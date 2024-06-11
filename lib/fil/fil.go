package fil

import (
	"coffee-monitor/lib/client"
	"coffee-monitor/lib/config"
	"coffee-monitor/lib/util"
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
)

func GetLotusInfo() (*LotusInfo, error) {
	ApiInit()
	conf := config.CONF
	chainHead, err := GetChainHead()
	if err != nil {
		return nil, err
	}
	height := gjson.Get(chainHead.Raw, "Height").Uint()
	log.Println("chainHead height:", height)

	peers, err := NetPeers()
	if err != nil {
		return nil, err
	}
	peersArray := peers.Array()
	log.Println("peers size:", len(peersArray))

	netAddrs, err := NetAddrsListen()
	if err != nil {
		return nil, err
	}
	id := gjson.Get(netAddrs.Raw, "ID").String()
	log.Println("ID:", id)

	return &LotusInfo{Id: id, Height: height, PeersNum: len(peersArray), Ip: util.GetLocalIP(), Account: conf.Fil.Account}, nil
}

func SendLotusInfo() error {
	lotusInfo, err := GetLotusInfo()
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(lotusInfo)
	if err != nil {
		return err
	}

	return client.SendData(string(jsonBytes))
}
