// BUG日志分析库文件
// 日志分析引擎核心程序，分析并记录BUG
// 使用正则匹配BUG关键词，并提取BUG内容/BUG数据汇总等

package lib

import (
	"bufio"
	config2 "coffee-monitor/lib/config"
	"coffee-monitor/lib/fil"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// 匹配关键词
func LogAnalysis(str string, keyword []string) (bool, string) {
	for _, v := range keyword {
		if strings.Contains(str, v) {
			return true, v
		}
	}
	return false, ""
}

// 分析一个日志文件
// 2024-05-06T10:57:40.175+0800	INFO	storageminer	storage/winning_prover.go:70	Computing WinningPoSt ;[{SealProof:9 SectorNumber:84016 SectorKey:<nil> SealedCID:bagboea4b5abcb6zhropclbjalxwnvfa7qi7ierozn4gxdpmsyo5zq6x4i4lhiyad}]; [103 44 215 64 74 116 106 107 26 18 126 176 236 251 158 238 155 38 167 255 96 112 136 73 159 55 54 126 226 10 67 231]
// 2024-05-06T10:57:40.175 INFO filecoin_proofs::api::post_util > generate_sector_challenges:start
// 2024-05-06T10:57:40.175 INFO filecoin_proofs::api::post_util > generate_sector_challenges:finish
// 2024-05-06T10:57:40.184 INFO filecoin_proofs::api::post_util > generate_single_vanilla_proof:start: SectorId(84016)
// 2024-05-06T10:57:40.382 INFO filecoin_proofs::api::post_util > generate_single_vanilla_proof:finish: SectorId(84016)
// 2024-05-06T10:57:41.489+0800	INFO	storageminer	storage/winning_prover.go:77	GenerateWinningPoSt took 1.313834905s
// 2024-05-06T10:57:42.034+0800	INFO	miner	miner/miner.go:660	mined new block	{"cid": "bafy2bzacecfv66423q22ciwve56jevpco7czuo2cldvfxqntq2zhm4sqwuoq6", "height": 3888596, "miner": "f02246008", "parents": ["f02245898","f01757676","f01366743","f02941888","f01082888","f01084913","f02830476","f01923787","f01964002","f02003555","f02182907"], "parentTipset": "{bafy2bzacedwdd6yur65buylz4536lhslwgljac6qmhbchf7f5hrlep7nwcf7m,bafy2bzaceacgnqni5rkbwgih7op6jr4zrythirm7h6rtetfnwxvdvo6iuc2os,bafy2bzaceb6btfwsuir3v7dgg64o3k723hfpdpepasc2vyzv4wplfgvtnhcd6,bafy2bzacea2tembbq7lj2osemaqisfr6cfet3gdg5tp6umkurzsti57suyq62,bafy2bzacebesgd7svgzwd6gqska5iw336tykhcv7ycmgsdb2cwqt7rs7ruhrs,bafy2bzacedvjaxacnfnngnc7xayus6wxw6ia3x2ngmrsbvyrxb2syq3jfnl5g,bafy2bzaceauvqhtkhu5myv6mwopt4nmvqq6d4uionexr7noqjqmnaxsat6mga,bafy2bzaceco6uvj52zgjzly4gj367vdjr54w22maipj2cuaeslipfzv72qc7a,bafy2bzacec4kzmo5aaoj37eqiyqjvpb4vmzf6oxsqzm32uudkduywl4r4vwkc,bafy2bzacec2dypzdmmgdzcryenickihx57giv4atng6s3rd26vy2bp4wflusi,bafy2bzacedn6ouurmcipctfsfgifzbjgztdmx7eikyoyt4v3dl7kyggrmdeac}", "took": 1.874394829}
func ReadNewBlocksFromLog(file string) ([]map[string]interface{}, error) {
	//文件操作
	var text []byte
	// 打开文件
	fs, err := Open(file)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs) // 读文件缓冲区

	ret := make([]map[string]interface{}, 0)
	//fmt.Print(ret)
	for io.EOF != err {
		text, _, err = buf.ReadLine() // 读一行
		if err == nil {
			block, err2 := ReadNewBlockFromLine(string(text))

			if err2 != nil {
				return nil, err2
			}

			if block != nil {
				ret = append(ret, block)
			}
		}
	}
	return ret, nil
}

func ReadNewBlockFromLine(content string) (map[string]interface{}, error) {
	var jsonIndex = strings.Index(content, "mined new block")
	if jsonIndex == -1 {
		return nil, nil
	}

	var timeIndex = strings.Index(content, "\tINFO\tminer")
	if timeIndex == -1 {
		return nil, nil
	}
	timeString := content[:timeIndex]
	jsonString := content[jsonIndex+15:]

	block, err := ParseJson(jsonString)

	if err != nil {
		return nil, err
	}
	block["time"] = timeString
	return block, nil
}

// 分析一个日志文件
func AnalysisLog(logPath string) {
	diffTime := time.Now().UnixNano()
	defer func() {
		// 显示程序执行效率
		diffTime = (time.Now().UnixNano() - diffTime) / 1e6
		fmt.Printf("程序共执行 %v ms \n", diffTime)
	}()

	fil.ApiInit()
	// 读配置文件
	config := config2.GetConfig()
	//fmt.Printf("%v\n\n", config)

	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Printf("处理的日期为：%v \n", date)

	for _, fileLogs := range config.Logfile {
		if len(logPath) > 0 {
			fileLogs.Path = logPath
		}
		blocks, err := ReadNewBlocksFromLog(fileLogs.Path)
		if err != nil {
			log.Printf("err:%v", err)
			return
		}
		log.Printf("start filter invalid block...")
		forkedNum := 0
		for _, block := range blocks {
			_, err := fil.GetBlock(block["cid"].(string))
			if err != nil {
				if strings.Contains(err.Error(), "ipld: could not find") {
					forkedNum++
					fmt.Printf("fored block time:%s, cid:%s, height:%f, took:%f \n", block["time"], block["cid"], block["height"], block["took"])
				} else {
					log.Printf("err:%v", err)
				}
			}
		}

		totalNum := len(blocks)
		forkedRate := float64(0)
		if totalNum > 0 {
			forkedRate = float64(forkedNum) / float64(totalNum) * 100
		}

		log.Printf("total number:%d, forked number:%d, forked rate is %.3f%s", totalNum, forkedNum, forkedRate, "%")
	}
}
