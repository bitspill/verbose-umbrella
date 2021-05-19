package oip5

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/azer/logger"
	"github.com/bitspill/flod/chaincfg/chainhash"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/gorilla/mux"
	"github.com/oipwg/proto/go/pb_oip"
	"github.com/oipwg/proto/go/pb_oip5/pb_templates/livenet"

	"github.com/oipwg/oip/flo"
	"github.com/oipwg/oip/httpapi"
)

func init() {
	o5Router.HandleFunc("/location/request", handleLocationRequest).Queries("id", "{id:[a-fA-F0-9]+}", "terms", "{terms}")
	o5Router.HandleFunc("/location/proof", handleLocationProof).Queries("id", "{id:[a-fA-F0-9]+}", "terms", "{terms}")
}

const commercialContentTypeUrl = "type.googleapis.com/oipProto.templates.tmpl_D8D0F22C"

type LocationRequestResponse struct {
	ValidUntil int64  `json:"valid_until"`
	Id         string `json:"id"`
	Term       string `json:"term"`
	PreImage   string `json:"pre_image"`
}

type LocationProofRequest struct {
	ValidUntil     int64  `json:"valid_until"`
	Id             string `json:"id"`
	Term           string `json:"term"`
	PreImage       string `json:"pre_image"`
	Signature      string `json:"signature"`
	PaymentTxid    string `json:"payment_txid"`
	SigningAddress string `json:"signing_address"`
}

type LocationProofResponse struct {
	ValidUntil int64  `json:"valid_until"`
	Id         string `json:"id"`
	Term       string `json:"term"`
	Network    string `json:"network"`
	Location   string `json:"location"`
}

func handleLocationProof(w http.ResponseWriter, r *http.Request) {
	var opts = mux.Vars(r)
	termString := opts["terms"]

	proofPost := &LocationProofRequest{}
	err := json.NewDecoder(r.Body).Decode(proofPost)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unable to decode proof",
		})
		log.Error("unable decode proof", logger.Attrs{"err": err, "txid": opts["id"]})
		return
	}

	//ok, err := flo.CheckSignature(proofPost.SigningAddress, proofPost.Signature, proofPost.PreImage)
	//if err != nil {
	//	httpapi.RespondJSON(w, 400, map[string]interface{}{
	//		"error": "unable to validate signature",
	//	})
	//	log.Error("unable to validate signature", logger.Attrs{"err": err, "txid": opts["id"]})
	//	return
	//}
	//
	//if !ok {
	//	httpapi.RespondJSON(w, 400, map[string]interface{}{
	//		"error": "invalid signature",
	//	})
	//	log.Error("invalid signature", logger.Attrs{"err": err, "txid": opts["id"]})
	//	return
	//}

	rec, err := GetRecord(opts["id"])
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unable to get record",
		})
		log.Error("unable to get record", logger.Attrs{"err": err, "txid": opts["id"]})
		return
	}

	comCont, term, err := getTerm(rec, termString)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": err,
		})
		log.Error("unable to obtain terms", logger.Attrs{"err": err})
		return
	}

	scs, ok := term.(*livenet.SimpleCoinSale)
	if !ok {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unsupported term",
		})
		log.Error("unsupported term type", logger.Attrs{"term": term})
		return
	}

	txh := &chainhash.Hash{}
	err = chainhash.Decode(txh, proofPost.PaymentTxid)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "invalid txid",
		})
		log.Error("unable to decode txid", logger.Attrs{"txid": opts["id"]})
		return
	}

	paid := false
	if bytes.Equal(scs.Coin.Raw, []byte("f9964d1e840608b68a3795fd2597e9b232dfce1029251d481b2110c83a68adf7")) {
		paid, err = checkFloPayment(txh, scs, proofPost.SigningAddress)
	}

	if bytes.Equal(scs.Coin.Raw, []byte("0000000000000000000000000000000000000000000000000000000000000001")) {
		paid, err = checkRvnPayment(txh, scs, proofPost.SigningAddress)
	}
	if bytes.Equal(scs.Coin.Raw, []byte("0000000000000000000000000000000000000000000000000000000000000002")) {
		paid, err = checkRvnAsset(scs, proofPost.SigningAddress)
	}

	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unable to check proof",
		})
		log.Error("unable to check payment", logger.Attrs{"paymentId": proofPost.PaymentTxid, "signingAddress": proofPost.SigningAddress, "id": opts["id"], "term": termString, "err": err})
		return
	}

	if !paid {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "insufficient proof",
		})
		log.Error("insufficient proof", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString})
		return
	}

	t := time.Now()

	httpapi.RespondJSON(w, 200, LocationProofResponse{
		ValidUntil: t.Add(3 * time.Minute).Unix(),
		Id:         opts["id"],
		Term:       termString,
		Network:    comCont.Network.String(),
		Location:   comCont.Location,
	})
}

func checkFloPayment(txh *chainhash.Hash, scs *livenet.SimpleCoinSale, signingAddress string) (bool, error) {
	tx, err := flo.GetTxVerbose(txh)
	if err != nil {
		return false, fmt.Errorf("unable to optain vinTx: %w", err)
	}

	paid := false
voutLoop:
	for _, vout := range tx.Vout {
		if vout.Value >= float64(scs.Amount)/float64(scs.Scale) {
			for _, address := range vout.ScriptPubKey.Addresses {
				if address == scs.Destination {
					paid = true
					break voutLoop
				}
			}
		}
	}

	if !paid {
		return false, nil
	}

	vinTxh := &chainhash.Hash{}
	paid = false
vinLoop:
	for _, vin := range tx.Vin {
		err := chainhash.Decode(vinTxh, vin.Txid)
		if err != nil {
			// should be impossible...
			return false, fmt.Errorf("unable to decode vin txid: %w", err)
		}

		vinTx, err := flo.GetTxVerbose(vinTxh)
		if err != nil {
			//log.Error("unable to obtain vinTx", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString, "vinTxid": vin.Txid})
			return false, fmt.Errorf("unable to obtain vinTx: %w", err)
		}

		for _, vout := range vinTx.Vout {
			for _, address := range vout.ScriptPubKey.Addresses {
				if address == signingAddress {
					paid = true
					break vinLoop
				}
			}
		}
	}
	return paid, nil
}

func checkRvnPayment(txh *chainhash.Hash, scs *livenet.SimpleCoinSale, signingAddress string) (bool, error) {
	paid := false

	txid := txh.String()
	res, err := http.Get("https://explorer-api.ravenland.org/tx/" + txid)
	if err != nil {
		return false, fmt.Errorf("tx request failed: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return false, fmt.Errorf("tx read failed: %w", err)
	}

	var txResponse RavenLandApiTxResponse
	err = json.Unmarshal(body, &txResponse)
	if err != nil {
		return false, fmt.Errorf("tx decode failed: %w", err)
	}

	for _, vout := range txResponse.Tx.Vout {
		if vout.Quantity >= float64(scs.Amount)/float64(scs.Scale) && vout.Asset == "RVN" {
			if vout.Address == scs.Destination {
				paid = true
				break
			}
		}
	}

	if !paid {
		return false, nil
	}

	// reset state variable to check sender
	paid = false

	for _, vin := range txResponse.Tx.Vin {
		if vin.VinType == "basic" && vin.Address == signingAddress {
			paid = true
			break
		}
	}

	return paid, nil
}

func checkRvnAsset(scs *livenet.SimpleCoinSale, signingAddress string) (bool, error) {
	paid := false

	res, err := http.Get("https://explorer-api.ravenland.org/address/" + signingAddress + "/balances")
	if err != nil {
		return false, fmt.Errorf("balance request failed: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return false, fmt.Errorf("balance read failed: %w", err)
	}

	var balResponse RavenLandApiBalanceResponse
	err = json.Unmarshal(body, &balResponse)
	if err != nil {
		return false, fmt.Errorf("balance decode failed: %w", err)
	}

	bal, ok := balResponse.Data[scs.Destination]
	if !ok {
		return false, nil
	}

	paid = bal > float64(scs.Amount)/float64(scs.Scale)

	return paid, nil
}

func handleLocationRequest(w http.ResponseWriter, r *http.Request) {
	var opts = mux.Vars(r)

	termString := opts["terms"]

	rec, err := GetRecord(opts["id"])
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unable to get record",
		})
		log.Error("unable to get record", logger.Attrs{"err": err, "txid": opts["id"]})
		return
	}

	_, term, err := getTerm(rec, termString)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": err,
		})
		log.Error("unable to obtain terms", logger.Attrs{"err": err})
		return
	}

	scs, ok := term.(*livenet.SimpleCoinSale)
	if !ok {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unsupported term",
		})
		log.Error("unsupported term type", logger.Attrs{"term": term})
		return
	}

	t := time.Now()

	httpapi.RespondJSON(w, 200, LocationRequestResponse{
		Id:         opts["id"],
		Term:       termString,
		PreImage:   fmt.Sprintf("%s-%s-%d-%d", opts["id"], termString, scs.Amount, t.Unix()),
		ValidUntil: t.Add(3 * time.Minute).Unix(),
	})
}

func getTerm(rec *oip5Record, termString string) (*livenet.CommercialContent, interface{}, error) {
	if strings.ToLower(termString) != "3733247363" {
		return nil, nil, errors.New("only simple coin sale is supported")
	}

	comCont, err := getCommercialContent(rec)
	if err != nil {
		return nil, nil, err
	}

	if comCont == nil {
		log.Error("commercial content information not found")
		return nil, nil, errors.New("commercial content information not found")
	}

	if len(termString) == 64 {
		t, err := getExternalTerms(rec, comCont, termString)
		return comCont, t, err
	}
	t, err := getEmbeddedTerm(rec, comCont, termString)
	return comCont, t, err
}

func getEmbeddedTerm(rec *oip5Record, comCont *livenet.CommercialContent, term string) (interface{}, error) {
	emTerm, err := strconv.ParseUint(term, 10, 32)
	if err != nil {
		log.Error("invalid term argument")
		return nil, errors.New("invalid term argument")
	}
	found := false
	for _, t := range comCont.EmbeddedTerms {
		if uint32(emTerm) == t {
			found = true
			break
		}
	}
	if !found {
		log.Error("requested terms not specified in this record")
		return nil, errors.New("requested terms not specified in this record")
	}

	termTypeUrl := fmt.Sprintf("type.googleapis.com/oipProto.templates.tmpl_%08X", emTerm)

	var a *any.Any
	for _, det := range rec.Record.Details.Details {
		if det.TypeUrl == termTypeUrl {
			a = det
			break
		}
	}

	if emTerm == 3733247363 {
		scs := &livenet.SimpleCoinSale{}
		err := ptypes.UnmarshalAny(a, scs)
		if err != nil {
			return nil, errors.New("unable to unmarshal terms")
		}
		return scs, nil
	}

	return nil, errors.New("embedded term not found")
}

func getExternalTerms(rec *oip5Record, comCont *livenet.CommercialContent, term string) (interface{}, error) {
	found := false
	for _, t := range comCont.Terms {
		if term == pb_oip.TxidToString(t) {
			found = true
			break
		}
	}
	if !found {
		log.Error("requested terms not specified in this record")
		return nil, errors.New("requested terms not specified in this record")
	}
	log.Error("external terms unsupported")
	return nil, errors.New("external terms unsupported")
}

func getCommercialContent(rec *oip5Record) (*livenet.CommercialContent, error) {
	var comCont *livenet.CommercialContent
	for i := range rec.Record.Details.Details {
		if len(rec.Record.Details.Details[i].TypeUrl) == 52 &&
			rec.Record.Details.Details[i].TypeUrl[44:] == commercialContentTypeUrl[44:] {
			comCont = &livenet.CommercialContent{}
			err := ptypes.UnmarshalAny(rec.Record.Details.Details[i], comCont)
			if err != nil {
				log.Error("unable to unmarshal commercial content")
				return nil, errors.New("unable to unmarshal commercial content")
			}
			break
		}
	}
	return comCont, nil
}

type RavenLandApiTxResponse struct {
	Height int64              `json:"height"`
	Tx     RavenLandApiTxInfo `json:"tx"`
}

type RavenLandApiTxInfo struct {
	Txid        string             `json:"txid"`
	Hash        string             `json:"hash"`
	Time        int64              `json:"time"`
	Vin         []RavenLandApiVin  `json:"vin"`
	Vout        []RavenLandApiVout `json:"vout"`
	Size        int64              `json:"size"`
	Vsize       int64              `json:"vsize"`
	Locktime    int64              `json:"locktime"`
	BlockHeight int64              `json:"blockHeight"`
}

type RavenLandApiVin struct {
	VinType  string  `json:"type"`
	Quantity float64 `json:"quantity"`
	Address  string  `json:"address"`
	Asset    string  `json:"asset"`
}

type RavenLandApiVout struct {
	Address  string  `json:"address"`
	Quantity float64 `json:"quantity"`
	Asset    string  `json:"asset"`
}

type RavenLandApiBalanceResponse struct {
	Data      map[string]float64 `json:"data"`
	UpdatedAt int64              `json:"updatedAt"`
}
