package floTorrent

import (
	"encoding/json"
	"strings"

	"github.com/azer/logger"
	"github.com/oipwg/oip/config"
	"github.com/oipwg/oip/datastore"
	"github.com/oipwg/oip/events"
	"gopkg.in/olivere/elastic.v6"
)

func init() {
	log.Info("init FloTorrent")
	if !config.IsTestnet() {
		events.SubscribeAsync("flo:floData", floDataProcessor, false)
		events.SubscribeAsync("modules:floTorrent:torrent", onTorrent, false)
		events.SubscribeAsync("modules:floTorrent:data", onData, false)
		datastore.RegisterMapping("flo_torrent", "floTorrent.json")
	}
}

func floDataProcessor(floData string, tx *datastore.TransactionData) {
	if tx.Block < 3443500 {
		return
	}

	if strings.HasPrefix(floData, `{"FLO_Torrent"`) {
		events.Publish("modules:floTorrent:torrent", floData, tx)
		return
	}
	if strings.HasPrefix(floData, `{"data":"`) {
		events.Publish("modules:floTorrent:data", floData, tx)
		return
	}

}

func onTorrent(floData string, tx *datastore.TransactionData) {
	ftw := &floTorrentWraper{}
	err := json.Unmarshal([]byte(floData), ftw)
	if err != nil {
		log.Error("unable to unmarshal flodata", logger.Attrs{"err": err, "txid": tx.Transaction.Txid})
	}

	elf := elFloTorrent{
		FloTorrent: &ftw.FloTorrent,
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index("flo_torrent")).Type("_doc").Id(tx.Transaction.Txid).Doc(elf)
	datastore.AutoBulk.Add(bir)
}

func onData(floData string, tx *datastore.TransactionData) {
	ftd := &floTorrentData{}
	err := json.Unmarshal([]byte(floData), ftd)
	if err != nil {
		log.Error("unable to unmarshal flodata", logger.Attrs{"err": err, "txid": tx.Transaction.Txid})
	}

	elf := elFloTorrent{
		Data: ftd,
	}

	bir := elastic.NewBulkIndexRequest().Index(datastore.Index("flo_torrent")).Type("_doc").Id(tx.Transaction.Txid).Doc(elf)
	datastore.AutoBulk.Add(bir)
}

type elFloTorrent struct {
	FloTorrent *floTorrent     `json:"flo_torrent"`
	Data       *floTorrentData `json:"data"`
}

type floTorrentWraper struct {
	FloTorrent floTorrent `json:"FLO_Torrent"`
}

type floTorrent struct {
	Name        string `json:"name"`
	Filename    string `json:"filename"`
	Type        string `json:"type"`
	Size        int64  `json:"size"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Chunks      int64  `json:"chunks"`
	StartTx     string `json:"startTx"`
}

type floTorrentData struct {
	Data string  `json:"data"`
	Next ftdNext `json:"next"`
}

type ftdNext string

func (f *ftdNext) UnmarshalJSON(b []byte) error {
	if len(b) == 5 && string(b) == "false" {
		*f = ""
		return nil
	}
	var s string
	err := json.Unmarshal(b, &s)
	*f = ftdNext(s)
	return err
}
