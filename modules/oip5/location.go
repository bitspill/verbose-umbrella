package oip5

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/oipwg/oip/rvn"
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

var decodedFloBytes []byte
var decodedRvnBytes []byte

func init() {
	o5Router.HandleFunc("/location/request", handleLocationRequest).Queries("id", "{id:[a-fA-F0-9]+}", "terms", "{terms}")
	o5Router.HandleFunc("/location/proof", handleLocationProof).Queries("id", "{id:[a-fA-F0-9]+}", "terms", "{terms}")

	decodedFloBytes, _ = hex.DecodeString("f9964d1e840608b68a3795fd2597e9b232dfce1029251d481b2110c83a68adf7")
	decodedRvnBytes, _ = hex.DecodeString("e1a9b68de823bdee67f1a8c55167cc15bf7006129d9acf072721a4043932b389")
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

	paid := false
	switch termString {
	case "3733247363": // SimpleCoinSale
		scs, ok := term.(*livenet.SimpleCoinSale)
		if ok {
			done := false
			err, paid, done = simpleSale(w, proofPost, opts, scs, termString)
			if done {
				return
			}
		}
	case "3993842283": // SimpleAssetHeld
		sah, ok := term.(*livenet.SimpleAssetHeld)
		if ok {
			done := false
			err, paid, done = simpleAsset(w, proofPost, opts, sah, termString)
			if done {
				return
			}
		}
	default:
		// unsupported term
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

func simpleSale(w http.ResponseWriter, proofPost *LocationProofRequest, opts map[string]string, scs *livenet.SimpleCoinSale, termString string) (error, bool, bool) {
	txh := &chainhash.Hash{}
	err := chainhash.Decode(txh, proofPost.PaymentTxid)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "invalid txid",
		})
		log.Error("unable to decode txid", logger.Attrs{"txid": opts["id"]})
		return nil, false, true
	}
	ok := false
	paid := false

	switch {
	case bytes.Equal(scs.Coin.Raw, decodedFloBytes):
		log.Info("checking flo")
		ok, err = flo.CheckSignature(proofPost.SigningAddress, proofPost.Signature, proofPost.PreImage)
		if !ok || err != nil {
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "invalid signature",
			})
			log.Error("invalid signature", logger.Attrs{"err": err, "txid": opts["id"]})
			return nil, false, true
		}
		paid, err = checkFloPayment(txh, scs, proofPost.SigningAddress)

	case bytes.Equal(scs.Coin.Raw, decodedRvnBytes):
		log.Info("checking rvn")
		ok, err = rvn.CheckSignature(proofPost.SigningAddress, proofPost.Signature, proofPost.PreImage)
		if !ok || err != nil {
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "invalid signature",
			})
			log.Error("invalid signature", logger.Attrs{"err": err, "txid": opts["id"]})
			return nil, false, true
		}
		paid, err = checkRvnPayment(txh, scs, proofPost.SigningAddress)
	default:
		log.Error("No matching coin", logger.Attrs{"coin": scs.Coin.String(), "id": opts["id"], "term": termString})
	}
	return err, paid, false
}

func simpleAsset(w http.ResponseWriter, proofPost *LocationProofRequest, opts map[string]string, sah *livenet.SimpleAssetHeld, termString string) (error, bool, bool) {
	paid := false
	switch {
	case bytes.Equal(sah.Coin.Raw, decodedRvnBytes):
		log.Info("checking rvn asset")
		ok, err := rvn.CheckSignature(proofPost.SigningAddress, proofPost.Signature, proofPost.PreImage)
		if !ok || err != nil {
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "invalid signature",
			})
			log.Error("invalid signature", logger.Attrs{"err": err, "txid": opts["id"]})
			return nil, false, true
		}
		paid, err = checkRvnAsset(sah, proofPost.SigningAddress)
		return err, paid, false

	default:
		log.Error("No matching coin", logger.Attrs{"coin": sah.Coin.String(), "id": opts["id"], "term": termString})
		return nil, false, false
	}
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
		log.Info("unable to find vout")
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

func checkRvnAsset(sah *livenet.SimpleAssetHeld, signingAddress string) (bool, error) {
	paid := false

	if sah.Expires != 0 && time.Now().Unix() > int64(sah.Expires) {
		return false, errors.New("asset access expired")
	}

	res, err := http.Get("https://explorer-api.ravenland.org/address/" + signingAddress + "/balances")
	if err != nil {
		log.Error("RavenLand balance failure", spew.Sdump(err))
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

	bal, ok := balResponse.Data[sah.Asset]
	if !ok {
		return false, nil
	}

	paid = bal >= float64(sah.Amount)/float64(sah.Scale)

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

	amount := uint32(0)
	switch termString {
	case "3733247363": // SimpleCoinSale
		scs, ok := term.(*livenet.SimpleCoinSale)
		if !ok {
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "term error",
			})
			log.Error("error casting term", logger.Attrs{"term": term})
			return
		}
		amount = scs.Amount
	case "3993842283": // SimpleAssetHeld
		sah, ok := term.(*livenet.SimpleAssetHeld)
		if !ok {
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "term error",
			})
			log.Error("error casting term", logger.Attrs{"term": term})
			return
		}
		amount = sah.Amount
	default:
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
		PreImage:   fmt.Sprintf("%s-%s-%d-%d", opts["id"], termString, amount, t.Unix()),
		ValidUntil: t.Add(3 * time.Minute).Unix(),
	})
}

func getTerm(rec *oip5Record, termString string) (*livenet.CommercialContent, interface{}, error) {
	if strings.ToLower(termString) != "3733247363" && strings.ToLower(termString) != "3993842283" {
		return nil, nil, errors.New("only simple terms are supported")
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

	if emTerm == 3993842283 {
		sah := &livenet.SimpleAssetHeld{}
		err := ptypes.UnmarshalAny(a, sah)
		if err != nil {
			return nil, errors.New("unable to unmarshal terms")
		}
		return sah, nil
	}

	return nil, errors.New("embedded term not found")
}

func getExternalTerms(_ *oip5Record, comCont *livenet.CommercialContent, term string) (interface{}, error) {
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
