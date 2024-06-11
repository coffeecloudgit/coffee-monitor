package client

import (
	"coffee-monitor/lib/util"
	"fmt"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	//Connect()
}

func TestSend(t *testing.T) {
	jsonString := "{\"cid\": \"bafy2bzacecfv66423q22ciwve56jevpco7czuo2cldvfxqntq2zhm4sqwuoq6\", \"height\": 3888596, \"miner\": \"f02246008\", \"parents\": [\"f02245898\",\"f01757676\",\"f01366743\",\"f02941888\",\"f01082888\",\"f01084913\",\"f02830476\",\"f01923787\",\"f01964002\",\"f02003555\",\"f02182907\"], \"parentTipset\": \"{bafy2bzacedwdd6yur65buylz4536lhslwgljac6qmhbchf7f5hrlep7nwcf7m,bafy2bzaceacgnqni5rkbwgih7op6jr4zrythirm7h6rtetfnwxvdvo6iuc2os,bafy2bzaceb6btfwsuir3v7dgg64o3k723hfpdpepasc2vyzv4wplfgvtnhcd6,bafy2bzacea2tembbq7lj2osemaqisfr6cfet3gdg5tp6umkurzsti57suyq62,bafy2bzacebesgd7svgzwd6gqska5iw336tykhcv7ycmgsdb2cwqt7rs7ruhrs,bafy2bzacedvjaxacnfnngnc7xayus6wxw6ia3x2ngmrsbvyrxb2syq3jfnl5g,bafy2bzaceauvqhtkhu5myv6mwopt4nmvqq6d4uionexr7noqjqmnaxsat6mga,bafy2bzaceco6uvj52zgjzly4gj367vdjr54w22maipj2cuaeslipfzv72qc7a,bafy2bzacec4kzmo5aaoj37eqiyqjvpb4vmzf6oxsqzm32uudkduywl4r4vwkc,bafy2bzacec2dypzdmmgdzcryenickihx57giv4atng6s3rd26vy2bp4wflusi,bafy2bzacedn6ouurmcipctfsfgifzbjgztdmx7eikyoyt4v3dl7kyggrmdeac}\", \"took\": 1.874394829}"
	block, err := util.ParseJson(jsonString)

	msg := Message{Type: NEW_BLOCK, Data: block}

	if err != nil {
		fmt.Println(err)
		return
	}

	go ConnectServer()
	time.Sleep(2000 * time.Millisecond)
	err2 := SendMessage(msg)
	if err2 != nil {
		return
	}
	time.Sleep(5000 * time.Millisecond)
	DisConnectServer()

}
