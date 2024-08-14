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

func TestSendNewMineOne(t *testing.T) {
	jsonString := "{\"epoch\": 3888596, \"miner\": \"f02246008\"}"
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

func TestSendLotusMinerInfo(t *testing.T) {
	infoString := "Enabled subsystems: [Mining Sealing SectorStorage]\nStartTime: 314h41m30s (started at 2024-07-10 07:52:23 +0800 CST)\nChain: [sync ok] [basefee 100 aFIL]\n⚠ 1 Active alerts (check lotus-miner log alerts)\nMiner: f05xx8960 (32 GiB sectors)\nPower: 9.91 Pi / 22.7 Ei (0.0427%)\n        Raw: 1015 TiB / 5.662 EiB (0.0171%)\n        Committed: 1.063 PiB\n        Proving: 1015 TiB\nProjected average block win rate: 42.98/week (every 3h54m31s)\nProjected block win with 99.9% probability every 26h58m17s\n(projections DO NOT account for future network and miner growth)\n\nMiner Balance:    55398.537 FIL\n      PreCommit:  7.056 FIL\n      Pledge:     52371.423 FIL\n      Vesting:    155.325 FIL\n      Available:  2864.733 FIL\nMarket Balance:   255 FIL\n       Locked:    231.44 FIL\n       Available: 23.56 FIL\nWorker Balance:   34.359 FIL\n       Control:   88.59 FIL\nTotal Spendable:  3011.242 FIL\n\nBeneficiary:    f03155156\n\nSectors:\n        Total: 35660\n        Proving: 34898\n        AddPiece: 14\n        PreCommit1: 474\n        PreCommit2: 70\n        PreCommitBatchWait: 6\n        WaitSeed: 169\n        Committing: 11\n        CommitAggregateWait: 9\n        FinalizeSector: 7\n        Removed: 2\n\nWorkers: Seal(77) WdPoSt(1) WinPoSt(0)"

	msg := Message{Type: LotusMinerInfo, Content: infoString}

	go ConnectServer()
	time.Sleep(2000 * time.Millisecond)
	err2 := SendMessage(msg)
	if err2 != nil {
		return
	}
	time.Sleep(5000 * time.Millisecond)
	DisConnectServer()

}

//
