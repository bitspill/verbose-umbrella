package oip042

import (
	"github.com/azer/logger"
	"github.com/json-iterator/go"
	"github.com/oipwg/oip/datastore"
	"github.com/olivere/elastic/v7"
)

func on42JsonRegisterInfluencer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonRegisterInfluencer", logger.Attrs{"txid": tx.Transaction.Txid})

	sig := any.Get("signature").ToString()

	var el elasticOip042Influencer
	el.Influencer = any.GetInterface()
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

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042InfluencerIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonEditInfluencer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonEditInfluencer", logger.Attrs{"txid": tx.Transaction.Txid})

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
		Type:      "influencer",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042EditIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonTransferInfluencer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonTransferInfluencer", logger.Attrs{"txid": tx.Transaction.Txid})

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
		Type:      "influencer",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042TransferIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}

func on42JsonDeactivateInfluencer(any jsoniter.Any, tx *datastore.TransactionData) {
	t := log.Timer()
	defer t.End("on42JsonDeactivateInfluencer", logger.Attrs{"txid": tx.Transaction.Txid})

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
		Type:      "influencer",
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index(oip042DeactivateIndex)).Id(tx.Transaction.Txid).Doc(el)
	datastore.AutoBulk.Add(bir)
}
