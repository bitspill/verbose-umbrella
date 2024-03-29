package oip042

import (
	"github.com/azer/logger"
	"github.com/json-iterator/go"
	"github.com/oipwg/oip/datastore"
	"github.com/olivere/elastic/v7"
)

func on42JsonRegisterPub(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonRegisterPub", logger.Attrs{"txid": tx.Transaction.Txid})

	// name := any.Get("name").ToString()
	// if len(name) == 0 {
	//	log.Println("oip042 no pub.name")
	//	return
	// }

	sig := any.Get("signature").ToString()

	var el elasticOip042Pub
	el.Pub = any.GetInterface()
	el.Meta = AMeta{
		Block:       tx.Block,
		BlockHash:   tx.BlockHash,
		Deactivated: false,
		Signature:   sig,
		Time:        tx.Transaction.Time,
		Tx:          tx,
		Txid:        tx.Transaction.Txid,
		Type:        "oip042",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042PublisherIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonEditPub(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonEditPub", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Edit
	el.Edit = any.GetInterface()
	el.Meta = OMeta{
		Block:     tx.Block,
		BlockHash: tx.BlockHash,
		Completed: false,
		Signature: sig,
		Time:      tx.Transaction.Time,
		Tx:        tx,
		Txid:      tx.Transaction.Txid,
		Type:      "pub",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042EditIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonTransferPub(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonTransferPub", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Edit
	el.Edit = any.GetInterface()
	el.Meta = OMeta{
		Block:     tx.Block,
		BlockHash: tx.BlockHash,
		Completed: false,
		Signature: sig,
		Time:      tx.Transaction.Time,
		Tx:        tx,
		Txid:      tx.Transaction.Txid,
		Type:      "pub",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042TransferIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonDeactivatePub(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonDeactivatePub", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042DeactivateInterface
	el.Deactivate = any.GetInterface()
	el.Meta = OMeta{
		Block:     tx.Block,
		BlockHash: tx.BlockHash,
		Completed: false,
		Signature: sig,
		Time:      tx.Transaction.Time,
		Tx:        tx,
		Txid:      tx.Transaction.Txid,
		Type:      "pub",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042DeactivateIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}
