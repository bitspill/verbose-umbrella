package oip042

import (
	"github.com/azer/logger"
	"github.com/json-iterator/go"
	"github.com/oipwg/oip/datastore"
	"github.com/olivere/elastic/v7"
)

func on42JsonRegisterAutominer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonRegisterAutominer", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Autominer
	el.Autominer = any.GetInterface()
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

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042AutominerIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonEditAutominer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonEditAutominer", logger.Attrs{"txid": tx.Transaction.Txid})

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
		Type:      "autominer",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042EditIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonTransferAutominer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonTransferAutominer", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Transfer
	el.Transfer = any.GetInterface()
	el.Meta = OMeta{
		Block:     tx.Block,
		BlockHash: tx.BlockHash,
		Completed: false,
		Signature: sig,
		Time:      tx.Transaction.Time,
		Tx:        tx,
		Txid:      tx.Transaction.Txid,
		Type:      "autominer",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042TransferIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonDeactivateAutominer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonDeactivateAutominer", logger.Attrs{"txid": tx.Transaction.Txid})

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
		Type:      "autominer",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042DeactivateIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}
