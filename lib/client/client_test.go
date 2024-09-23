package client

import (
	_ "coffee-monitor/lib/log"
	"coffee-monitor/lib/util"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	//Connect()
}

func TestSendNewMineOne(t *testing.T) {
	jsonString := "{\"epoch\": \"3888596\", \"miner\": \"f02246008\"}"
	mineOne, err := util.ParseJson(jsonString)

	msg := Message{Type: NewMineOne, Data: mineOne}

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
func TestSendNewBlock(t *testing.T) {
	jsonString := "{\"cid\": \"bafy2bzacecfv66423q22ciwve56jevpco7czuo2cldvfxqntq2zhm4sqwuoq6\", \"height\": 3888596, \"miner\": \"f02246008\"," +
		" \"parents\": [\"f02245898\",\"f01757676\",\"f01366743\",\"f02941888\",\"f01082888\",\"f01084913\",\"f02830476\",\"f01923787\",\"f01964002\",\"f02003555\",\"f02182907\"], \"parentTipset\": \"{bafy2bzacedwdd6yur65buylz4536lhslwgljac6qmhbchf7f5hrlep7nwcf7m,bafy2bzaceacgnqni5rkbwgih7op6jr4zrythirm7h6rtetfnwxvdvo6iuc2os,bafy2bzaceb6btfwsuir3v7dgg64o3k723hfpdpepasc2vyzv4wplfgvtnhcd6,bafy2bzacea2tembbq7lj2osemaqisfr6cfet3gdg5tp6umkurzsti57suyq62,bafy2bzacebesgd7svgzwd6gqska5iw336tykhcv7ycmgsdb2cwqt7rs7ruhrs,bafy2bzacedvjaxacnfnngnc7xayus6wxw6ia3x2ngmrsbvyrxb2syq3jfnl5g,bafy2bzaceauvqhtkhu5myv6mwopt4nmvqq6d4uionexr7noqjqmnaxsat6mga,bafy2bzaceco6uvj52zgjzly4gj367vdjr54w22maipj2cuaeslipfzv72qc7a,bafy2bzacec4kzmo5aaoj37eqiyqjvpb4vmzf6oxsqzm32uudkduywl4r4vwkc,bafy2bzacec2dypzdmmgdzcryenickihx57giv4atng6s3rd26vy2bp4wflusi,bafy2bzacedn6ouurmcipctfsfgifzbjgztdmx7eikyoyt4v3dl7kyggrmdeac}\", \"took\": 1.874394829, " +
		"\"time\": \"2024-08-14T13:48:42.574+0800\", \"reward\": \"7.243937889909999 FIL\"}"
	block, err := util.ParseJson(jsonString)

	msg := Message{Type: NewBlock, Data: block}

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

func TestSendOrphanBlock(t *testing.T) {
	jsonString := "{\"cid\": \"bafy2bzacecfv66423q22ciwve56jevpco7czuo2cldvfxqntq2zhm4sqwuoq6\", \"height\": 3888596, \"miner\": \"f02246008\", \"parents\": [\"f02245898\",\"f01757676\",\"f01366743\",\"f02941888\",\"f01082888\",\"f01084913\",\"f02830476\",\"f01923787\",\"f01964002\",\"f02003555\",\"f02182907\"], \"parentTipset\": \"{bafy2bzacedwdd6yur65buylz4536lhslwgljac6qmhbchf7f5hrlep7nwcf7m,bafy2bzaceacgnqni5rkbwgih7op6jr4zrythirm7h6rtetfnwxvdvo6iuc2os,bafy2bzaceb6btfwsuir3v7dgg64o3k723hfpdpepasc2vyzv4wplfgvtnhcd6,bafy2bzacea2tembbq7lj2osemaqisfr6cfet3gdg5tp6umkurzsti57suyq62,bafy2bzacebesgd7svgzwd6gqska5iw336tykhcv7ycmgsdb2cwqt7rs7ruhrs,bafy2bzacedvjaxacnfnngnc7xayus6wxw6ia3x2ngmrsbvyrxb2syq3jfnl5g,bafy2bzaceauvqhtkhu5myv6mwopt4nmvqq6d4uionexr7noqjqmnaxsat6mga,bafy2bzaceco6uvj52zgjzly4gj367vdjr54w22maipj2cuaeslipfzv72qc7a,bafy2bzacec4kzmo5aaoj37eqiyqjvpb4vmzf6oxsqzm32uudkduywl4r4vwkc,bafy2bzacec2dypzdmmgdzcryenickihx57giv4atng6s3rd26vy2bp4wflusi,bafy2bzacedn6ouurmcipctfsfgifzbjgztdmx7eikyoyt4v3dl7kyggrmdeac}\", \"took\": 1.874394829}"
	block, err := util.ParseJson(jsonString)

	msg := Message{Type: OrphanBlock, Data: block}

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

func TestSendLotusMinerSectorsInfo(t *testing.T) {
	minerSectorInfoJsonString := "{\"miner\":\"f02246008\",\"sectors\":[{\"day\":\"2025-11-06\",\"from\":0,\"to\":7,\"sectorNum\":8},{\"day\":\"2025-11-07\",\"from\":8,\"to\":2916,\"sectorNum\":904},{\"day\":\"2025-11-08\",\"from\":1091,\"to\":2912,\"sectorNum\":31},{\"day\":\"2025-11-09\",\"from\":1354,\"to\":2982,\"sectorNum\":488},{\"day\":\"2025-11-10\",\"from\":2981,\"to\":21813,\"sectorNum\":1648},{\"day\":\"2025-11-11\",\"from\":4605,\"to\":22206,\"sectorNum\":2878},{\"day\":\"2025-11-12\",\"from\":7490,\"to\":10503,\"sectorNum\":2985},{\"day\":\"2025-11-13\",\"from\":10499,\"to\":13394,\"sectorNum\":2879},{\"day\":\"2025-11-14\",\"from\":13371,\"to\":22217,\"sectorNum\":3508},{\"day\":\"2025-11-15\",\"from\":16892,\"to\":22184,\"sectorNum\":2478},{\"day\":\"2025-11-16\",\"from\":19353,\"to\":22220,\"sectorNum\":2326},{\"day\":\"2025-11-17\",\"from\":22221,\"to\":25609,\"sectorNum\":2944},{\"day\":\"2025-11-18\",\"from\":24742,\"to\":29461,\"sectorNum\":3720},{\"day\":\"2025-11-19\",\"from\":28421,\"to\":30145,\"sectorNum\":1047},{\"day\":\"2025-11-20\",\"from\":29778,\"to\":31192,\"sectorNum\":1257},{\"day\":\"2025-11-21\",\"from\":31187,\"to\":31633,\"sectorNum\":443},{\"day\":\"2025-11-22\",\"from\":31634,\"to\":32105,\"sectorNum\":466},{\"day\":\"2025-11-23\",\"from\":32098,\"to\":33431,\"sectorNum\":1322},{\"day\":\"2025-11-24\",\"from\":33419,\"to\":46956,\"sectorNum\":2873},{\"day\":\"2025-11-25\",\"from\":36228,\"to\":37958,\"sectorNum\":1649},{\"day\":\"2025-11-26\",\"from\":37954,\"to\":40920,\"sectorNum\":2688},{\"day\":\"2025-11-27\",\"from\":40492,\"to\":44731,\"sectorNum\":3736},{\"day\":\"2025-11-28\",\"from\":44034,\"to\":54858,\"sectorNum\":3105},{\"day\":\"2025-11-29\",\"from\":46973,\"to\":55134,\"sectorNum\":4707},{\"day\":\"2025-11-30\",\"from\":51752,\"to\":55109,\"sectorNum\":634},{\"day\":\"2025-12-01\",\"from\":52896,\"to\":55659,\"sectorNum\":1903},{\"day\":\"2025-12-02\",\"from\":55348,\"to\":58835,\"sectorNum\":3156},{\"day\":\"2025-12-03\",\"from\":58396,\"to\":61020,\"sectorNum\":2411},{\"day\":\"2025-12-04\",\"from\":61021,\"to\":61923,\"sectorNum\":899},{\"day\":\"2025-12-05\",\"from\":61919,\"to\":63047,\"sectorNum\":1112},{\"day\":\"2025-12-06\",\"from\":63039,\"to\":65084,\"sectorNum\":2040},{\"day\":\"2025-12-07\",\"from\":65077,\"to\":68634,\"sectorNum\":1683},{\"day\":\"2025-12-08\",\"from\":65669,\"to\":68266,\"sectorNum\":1316},{\"day\":\"2025-12-09\",\"from\":67997,\"to\":68633,\"sectorNum\":548},{\"day\":\"2025-12-10\",\"from\":68635,\"to\":71381,\"sectorNum\":2735},{\"day\":\"2025-12-11\",\"from\":71365,\"to\":71937,\"sectorNum\":547}],\"total\":69074}"
	sectorsInfo, _ := util.ParseJson(minerSectorInfoJsonString)

	msg := Message{Type: SectorsExpireInfo, Data: sectorsInfo}

	go ConnectServer()
	time.Sleep(2000 * time.Millisecond)
	err2 := SendMessage(msg)
	if err2 != nil {
		return
	}
	time.Sleep(50000 * time.Millisecond)
	DisConnectServer()

}

func TestVersion(t *testing.T) {
	fmt.Println(runtime.Version())
}

//
