package fil

import (
	"testing"
)

func TestReadNewBlockFromLine(t *testing.T) {

	//content := "// 2024-05-06T10:57:42.034+0800\tINFO\tminer\tminer/miner.go:660\tmined new block\t{\"cid\": \"bafy2bzacecfv66423q22ciwve56jevpco7czuo2cldvfxqntq2zhm4sqwuoq6\", \"height\": 3888596, \"miner\": \"f02246008\", \"parents\": [\"f02245898\",\"f01757676\",\"f01366743\",\"f02941888\",\"f01082888\",\"f01084913\",\"f02830476\",\"f01923787\",\"f01964002\",\"f02003555\",\"f02182907\"], \"parentTipset\": \"{bafy2bzacedwdd6yur65buylz4536lhslwgljac6qmhbchf7f5hrlep7nwcf7m,bafy2bzaceacgnqni5rkbwgih7op6jr4zrythirm7h6rtetfnwxvdvo6iuc2os,bafy2bzaceb6btfwsuir3v7dgg64o3k723hfpdpepasc2vyzv4wplfgvtnhcd6,bafy2bzacea2tembbq7lj2osemaqisfr6cfet3gdg5tp6umkurzsti57suyq62,bafy2bzacebesgd7svgzwd6gqska5iw336tykhcv7ycmgsdb2cwqt7rs7ruhrs,bafy2bzacedvjaxacnfnngnc7xayus6wxw6ia3x2ngmrsbvyrxb2syq3jfnl5g,bafy2bzaceauvqhtkhu5myv6mwopt4nmvqq6d4uionexr7noqjqmnaxsat6mga,bafy2bzaceco6uvj52zgjzly4gj367vdjr54w22maipj2cuaeslipfzv72qc7a,bafy2bzacec4kzmo5aaoj37eqiyqjvpb4vmzf6oxsqzm32uudkduywl4r4vwkc,bafy2bzacec2dypzdmmgdzcryenickihx57giv4atng6s3rd26vy2bp4wflusi,bafy2bzacedn6ouurmcipctfsfgifzbjgztdmx7eikyoyt4v3dl7kyggrmdeac}\", \"took\": 1.874394829}\n"
	//_, err := ReadNewBlockFromLine(content)
	//if err != nil {
	//	t.Errorf("err:%v", err)
	//}

}

func TestReadNewBlocksFromLog(t *testing.T) {
	file := "/Users/fangyu/GolandProjects/coffee-monitor/logs/miner-last-500000.log"
	blocks, err := ReadNewBlocksFromLog(file)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	for _, block := range blocks {
		t.Logf("block time:%s, cid:%s, height:%f, took:%f", block["time"], block["cid"], block["height"], block["took"])
	}
}

func TestAnalysisLog(t *testing.T) {
	file := "/Users/fangyu/GolandProjects/coffee-monitor/logs/miner-last-500000.log"
	AnalysisLog(file)
}

func TestAnalysisTailLog(t *testing.T) {
	MinerLogTailProcessor()
}
