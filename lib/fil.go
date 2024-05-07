package lib

import (
	"coffee-monitor/lib/fil_rpc"
	"github.com/tidwall/gjson"
)

var walletClient = fil_rpc.Client{BaseURL: "https://api.node.glif.io/rpc/v0", Debug: true}

func FilInit() {
	config := GetConfig()
	if len(config.Fil.Url) == 0 {
		panic("error config fil")
	}
	walletClient = fil_rpc.Client{BaseURL: config.Fil.Url, Debug: config.Fil.Debug}
}

func GetBlock(cid string) (*gjson.Result, error) {
	//params := []interface{}{cid}
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
