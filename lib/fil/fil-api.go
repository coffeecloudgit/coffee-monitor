package fil

import (
	"coffee-monitor/lib/config"
	"coffee-monitor/lib/fil_rpc"
	"github.com/tidwall/gjson"
)

var walletClient = fil_rpc.Client{BaseURL: "https://api.node.glif.io/rpc/v0", Debug: true}

func ApiInit() {
	conf := config.GetConfig()
	if len(conf.Fil.Url) == 0 {
		panic("error config fil")
	}
	walletClient = fil_rpc.Client{BaseURL: conf.Fil.Url, Debug: conf.Fil.Debug}
}

func GetBlock(cid string) (*gjson.Result, error) {
	callMsg := map[string]interface{}{
		"/": cid,
	}
	params := []interface{}{callMsg}
	result, err := walletClient.Call("Filecoin.ChainGetBlock", params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetChainHead() (*gjson.Result, error) {
	result, err := walletClient.Call("Filecoin.ChainHead", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NetPeers() (*gjson.Result, error) {
	result, err := walletClient.Call("Filecoin.NetPeers", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NetAddrsListen() (*gjson.Result, error) {
	result, err := walletClient.Call("Filecoin.NetAddrsListen", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
