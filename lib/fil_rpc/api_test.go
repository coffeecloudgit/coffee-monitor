package fil_rpc

import (
	"fmt"
	"testing"
)

const (
	//baseURL = "http://127.0.0.1:1234/rpc/v0" //fil_test_local
	baseURL   = "http://47.57.26.144:20031/rpc/v0" //fil_test_remote
	CidLength = 62
)

func TestGetCall(t *testing.T) {
	client := Client{BaseURL: baseURL}
	method := "Filecoin.ChainHead"
	//method := "Filecoin.ChainGetTipSetByHeight"
	//method := "Filecoin.ChainGetBlock"
	//method := "Filecoin.ChainGetBlockMessages"
	//method := "Filecoin.WalletBalance"
	//method := "Filecoin.ChainGetTipSet"
	//method := "Filecoin.StateGetReceipt"
	//method := "Filecoin.MpoolGetNonce"
	//method := "Filecoin.MpoolEstimateGasPrice"
	//method := "Filecoin.StateGetActor"

	//blockCids := make([]interface{}, 0)
	//tipSetKey := []interface{}{ blockCids }

	params := []interface{}{
		//"t1abuzgc6y4tirvo274gyayqvslfmgkua3ksuyu4y",
		//blockCids,
	}

	//for i := 0; i <= 10; i++ {
	result, err := client.Call(method, params)
	if err != nil {
		t.Logf("Get Call Result return: \n\t%+v\n", err)
	}

	//for _, block := range gjson.Get(result.Raw, "Blocks").Array() {
	//	height := uint64( gjson.Get(block.Raw, "Height").Uint() )
	//	fmt.Println(", height:", height )
	//}

	if result != nil {
		fmt.Println(method, ", result:", result.String())
	}
	//}
}

func TestGetCallWithToken(t *testing.T) {

}

// return : []{"/":"ba..."}...
func GetTipSetKey(hash string) []interface{} {
	blockCids := make([]interface{}, 0)
	for i := 0; i < (len(hash) / CidLength); i++ {
		blockCidHash := hash[i*CidLength : (i+1)*CidLength]

		blockCid := map[string]interface{}{
			"/": blockCidHash,
		}

		blockCids = append(blockCids, blockCid)
	}
	return blockCids
}
