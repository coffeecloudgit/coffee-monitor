package fil

import (
	"coffee-monitor/lib/config"
	"coffee-monitor/lib/fil_rpc"
	"github.com/tidwall/gjson"
)

func init() {
	ApiInit()
}

var walletClient = fil_rpc.Client{BaseURL: "https://api.node.glif.io/rpc/v0", Debug: true}

func ApiInit() {
	conf := config.CONF
	if len(conf.Fil.Url) == 0 {
		panic("error config fil")
	}
	walletClient = fil_rpc.Client{BaseURL: conf.Fil.Url, Debug: conf.Fil.Debug}
}

// AV#TWVpT2!FjWYXm
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

func GetBlockMessage(cid string) (*gjson.Result, error) {
	callMsg := map[string]interface{}{
		"/": cid,
	}
	params := []interface{}{callMsg}
	result, err := walletClient.Call("Filecoin.ChainGetBlockMessages", params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetTipSetByHeight(height uint64) (*gjson.Result, error) {
	params := []interface{}{height, nil}

	result, err := walletClient.Call("Filecoin.ChainGetTipSetByHeight", params)
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

func SyncState() (*gjson.Result, error) {
	result, err := walletClient.Call("Filecoin.SyncState", nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
