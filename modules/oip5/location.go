package oip5

import (
	"encoding/json"
	"errors"
	"fmt"
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

	ok, err := flo.CheckSignature(proofPost.SigningAddress, proofPost.Signature, proofPost.PreImage)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unable to validate signature",
		})
		log.Error("unable to validate signature", logger.Attrs{"err": err, "txid": opts["id"]})
		return
	}

	if !ok {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "invalid signature",
		})
		log.Error("invalid signature", logger.Attrs{"err": err, "txid": opts["id"]})
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

	tx, err := flo.GetTxVerbose(txh)
	if err != nil {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "unable to locate transaction",
		})
		log.Error("unable to obtain vinTx", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString})
		return
	}

	// ToDo: it's valid to send multiple vout to the same destination
	//  but we'll only check for single vout covering full amount currently
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
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "insufficient payment",
		})
		log.Error("insufficient payment", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString})
		return
	}

	vinTxh := &chainhash.Hash{}
	paid = false
vinLoop:
	for _, vin := range tx.Vin {
		err := chainhash.Decode(vinTxh, vin.Txid)
		if err != nil {
			// should be impossible...
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "unable to decode transaction",
			})
			log.Error("unable to decode vin txid", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString, "vinTxid": vin.Txid})
			return
		}

		vinTx, err := flo.GetTxVerbose(vinTxh)
		if err != nil {
			httpapi.RespondJSON(w, 400, map[string]interface{}{
				"error": "unable to locate transaction",
			})
			log.Error("unable to obtain vinTx", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString, "vinTxid": vin.Txid})
			return
		}

		for _, vout := range vinTx.Vout {
			for _, address := range vout.ScriptPubKey.Addresses {
				if address == proofPost.SigningAddress {
					paid = true
					break vinLoop
				}
			}
		}
	}
	if !paid {
		httpapi.RespondJSON(w, 400, map[string]interface{}{
			"error": "insufficient payment",
		})
		log.Error("insufficient payment", logger.Attrs{"paymentId": proofPost.PaymentTxid, "id": opts["id"], "term": termString})
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
