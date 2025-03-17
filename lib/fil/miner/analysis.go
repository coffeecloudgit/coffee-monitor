// BUG日志分析库文件
// 日志分析引擎核心程序，分析并记录BUG
// 使用正则匹配BUG关键词，并提取BUG内容/BUG数据汇总等

package fil

import (
	"bufio"
	"coffee-monitor/lib/client"
	config2 "coffee-monitor/lib/config"
	"coffee-monitor/lib/fil"
	"coffee-monitor/lib/log"
	"coffee-monitor/lib/shell"
	"coffee-monitor/lib/util"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/hpcloud/tail"
	"github.com/tidwall/gjson"
)

// LogAnalysis 匹配关键词
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
	fs, err := util.Open(file)
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

func ReadNewMineOneFromLine(content string) (map[string]interface{}, error) {
	//2024-07-23T18:32:40.022+0800    INFO    miner   miner/miner.go:505      completed mineOne       {"tookMilliseconds": 10, "forRound": 4114146, "baseEpoch": 4114145, "baseDeltaSeconds": 10, "nullRounds": 0, "lateStart": false, "beaconEpoch": 9642462, "lookbackEpochs": 900, "networkPowerAtLookback": "26141019317397585920", "minerPowerAtLookback": "11157637840240640", "isEligible": true, "isWinner": false, "error": null}
	var timeIndex = strings.Index(content, "\tINFO\tminer")
	var timeString string
	var mineOne map[string]interface{}
	var err error
	var minerJsonIndex = strings.Index(content, "completed mineOne")

	if minerJsonIndex == -1 {
		return nil, nil
	}

	minerJsonString := content[minerJsonIndex+17:]
	mineOne, err = util.ParseJson(minerJsonString)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if timeIndex > -1 {
		timeString = content[:timeIndex]
		mineOne["time"] = timeString
	}

	if timeIndex == -1 {
		return nil, nil
	}
	epoch := mineOne["baseEpoch"]
	mineOne["baseEpoch"] = strconv.FormatFloat(epoch.(float64), 'f', 0, 64) //fmt.Sprintf("%f", epoch)
	return mineOne, nil
}

func ReadNewBlockFromLine(content string) (map[string]interface{}, error) {
	//2024-07-23T18:32:40.022+0800    INFO    miner   miner/miner.go:505      completed mineOne       {"tookMilliseconds": 10, "forRound": 4114146, "baseEpoch": 4114145, "baseDeltaSeconds": 10, "nullRounds": 0, "lateStart": false, "beaconEpoch": 9642462, "lookbackEpochs": 900, "networkPowerAtLookback": "26141019317397585920", "minerPowerAtLookback": "11157637840240640", "isEligible": true, "isWinner": false, "error": null}
	var timeIndex = strings.Index(content, "\tINFO\tminer")
	var timeString string
	var err error

	var jsonIndex = strings.Index(content, "mined new block")
	if jsonIndex == -1 {
		return nil, nil
	}

	if timeIndex == -1 {
		return nil, nil
	}
	timeString = content[:timeIndex]
	jsonString := content[jsonIndex+15:]

	block, err := util.ParseJson(jsonString)

	if err != nil {
		return nil, err
	}
	block["time"] = timeString
	block["reward"] = ""

	return block, nil
}

// AnalysisLog 分析一个日志文件
func AnalysisLog(logPath string) {
	diffTime := time.Now().UnixNano()
	defer func() {
		// 显示程序执行效率
		diffTime = (time.Now().UnixNano() - diffTime) / 1e6
		fmt.Printf("程序共执行 %v ms \n", diffTime)
	}()

	fil.ApiInit()
	// 读配置文件
	config := config2.CONF
	//fmt.Printf("%v\n\n", config)

	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Printf("处理的日期为：%v \n", date)

	for _, fileLogs := range config.Logfile {
		if len(logPath) > 0 {
			fileLogs.Path = logPath
		}
		blocks, err := ReadNewBlocksFromLog(fileLogs.Path)
		if err != nil {
			log.Logger.Info("err:%v", err)
			return
		}
		log.Logger.Info("start filter invalid block...")
		forkedNum := 0
		for _, block := range blocks {
			_, err := fil.GetBlock(block["cid"].(string))
			if err != nil {
				if strings.Contains(err.Error(), "ipld: could not find") {
					forkedNum++
					fmt.Printf("fored block time:%s, cid:%s, height:%f, took:%f, parents:%v \n",
						block["time"], block["cid"], block["height"], block["took"], block["parents"])

					//fmt.Printf("height:%f, paretns:%v => %v \n", block["height"], block["parents"], block["parents"])
				} else {
					log.Logger.Info("err:%v", err)
				}
			}
		}

		totalNum := len(blocks)
		forkedRate := float64(0)
		if totalNum > 0 {
			forkedRate = float64(forkedNum) / float64(totalNum) * 100
		}

		log.Logger.Info("total number:%d, forked number:%d, forked rate is %.3f%s", totalNum, forkedNum, forkedRate, "%")
	}
}

/*
*
{"cid": "bafy2bzacecfv66423q22ciwve56jevpco7czuo2cldvfxqntq2zhm4sqwuoq6",
"height": 3888596,
"miner": "f02246008",
"parents": ["f02245898","f01757676","f01366743","f02941888","f01082888","f01084913","f02830476","f01923787","f01964002","f02003555","f02182907"],
"parentTipset": "{bafy2bzacedwdd6yur65buylz4536lhslwgljac6qmhbchf7f5hrlep7nwcf7m,bafy2bzaceacgnqni5rkbwgih7op6jr4zrythirm7h6rtetfnwxvdvo6iuc2os,bafy2bzaceb6btfwsuir3v7dgg64o3k723hfpdpepasc2vyzv4wplfgvtnhcd6,bafy2bzacea2tembbq7lj2osemaqisfr6cfet3gdg5tp6umkurzsti57suyq62,bafy2bzacebesgd7svgzwd6gqska5iw336tykhcv7ycmgsdb2cwqt7rs7ruhrs,bafy2bzacedvjaxacnfnngnc7xayus6wxw6ia3x2ngmrsbvyrxb2syq3jfnl5g,bafy2bzaceauvqhtkhu5myv6mwopt4nmvqq6d4uionexr7noqjqmnaxsat6mga,bafy2bzaceco6uvj52zgjzly4gj367vdjr54w22maipj2cuaeslipfzv72qc7a,bafy2bzacec4kzmo5aaoj37eqiyqjvpb4vmzf6oxsqzm32uudkduywl4r4vwkc,bafy2bzacec2dypzdmmgdzcryenickihx57giv4atng6s3rd26vy2bp4wflusi,bafy2bzacedn6ouurmcipctfsfgifzbjgztdmx7eikyoyt4v3dl7kyggrmdeac}",
"took": 1.874394829}
*/
//var blockQueue = util.Queue{
//	Content: []util.Object{},
//	Timeout: 0,
//	MaxSize: 60,
//}

var blockQueue = make(map[string]map[string]interface{}, 5)

//var blockQueueLock sync.Mutex

func MinerLogTailProcessor() error {
	go client.ConnectServer()
	config := config2.CONF
	if len(config.Logfile) == 0 {
		return errors.New("miner logfile is empty")
	}
	//t, err := tail.TailFile(config.Logfile[0].Path, tail.Config{Follow: true})

	t, err := tail.TailFile(config.Logfile[0].Path, tail.Config{
		ReOpen:    true,                                              //是否重新打开
		Follow:    true,                                              //是否跟随
		Location:  &tail.SeekInfo{Offset: -5000, Whence: io.SeekEnd}, //从文件的什么地方开始读
		MustExist: false,                                             //文件不存在不报错
		Poll:      false,
	})

	if err != nil {
		log.Logger.Info(err.Error())
		return err
	}
	for line := range t.Lines {
		//log.Logger.Info(line.Text)
		mineOne, _ := ReadNewMineOneFromLine(line.Text)

		if mineOne != nil {
			mineOneTime := mineOne["time"]
			myTime, err3 := util.StringToTime(mineOneTime.(string))

			if err3 != nil {
				log.Logger.Info(err3.Error())
			}
			if err3 == nil {
				//log.Logger.Info(myTime)
				secondSub := time.Now().Unix() - myTime.Unix()

				if secondSub < 300 {
					mine := make(map[string]interface{}, 5)
					mine["epoch"] = mineOne["baseEpoch"]
					mine["miner"] = config.Fil.Account
					msg := client.Message{Type: client.NewMineOne, Data: mine}
					client.SendMessage(msg)
				}
			}

		}
		block, _ := ReadNewBlockFromLine(line.Text)
		if block != nil {
			//msg := client.Message{Type: client.NewBlock, Data: block}
			cid := block["cid"].(string)
			if cid == "" {
				continue
			}
			//err2, reward := shell.LotusMinerInfoGetRewardForBlock(cid)
			//if err2 != nil {
			//	fmt.Println(err2.Error())
			//}
			//block["reward"] = reward
			//client.SendMessage(msg)
			block["send"] = false
			blockQueue[cid] = block

		}
	}
	return nil
}

func CheckOrphanBlock() {
	log.Logger.Info("check orphan block:", slog.Int("queueLength", len(blockQueue)))
	var deleteKeys = make([]string, 0)
	if len(blockQueue) <= 0 {
		return
	}
	info, err := fil.GetLotusInfo()
	if err != nil {
		log.Logger.Info(err.Error())
		return
	}
	log.Logger.Info("start check orphan block, len:",
		slog.Int("queueLength", len(blockQueue)), slog.Uint64("height", info.Height))

	err, rewardMap := shell.LotusMinerInfoGetRewardForLastBlocks()
	if err != nil {
		log.Logger.Error(err.Error())
	}
	for cid, block := range blockQueue {
		height, err2 := util.InterfaceToUint64(block["height"])
		if err2 != nil {
			log.Logger.Info(err2.Error())
			continue
		}

		//一次确认后发送出块信息
		if !block["send"].(bool) && (info.Height-height) > 1 {
			log.Logger.Info("new block send, len:", slog.Int("queueLength", len(blockQueue)),
				slog.Uint64("info height", info.Height), slog.Uint64("height", height))
			// err3, reward := shell.LotusMinerInfoGetRewardForBlock(cid)
			// if err3 != nil {
			// 	log.Logger.Error(err3.Error())
			// 	continue
			// }

			reward, ok := rewardMap[cid]
			if ok {
				log.Logger.Info("reward found", "cid:", cid, "reward:", reward)
				//continue
				block["reward"] = reward
				msg := client.Message{Type: client.NewBlock, Data: block}
				client.SendMessage(msg)
				block["send"] = true
			}

		}

		if (info.Height - height) < 30 { //30个确认后判断孤块
			continue
		}

		if (info.Height - height) > 2000 { //2000个确认后不能判断孤块，直接删除
			deleteKeys = append(deleteKeys, cid)
			continue
		}
		parentHeight := height - 1
		tipSet, err2 := fil.GetTipSetByHeight(parentHeight)
		if err2 != nil {
			log.Logger.Error("err:", "err", err, "block", block)
		}
		//log.Logger.Info("tipSet", "tipSet", tipSet)
		cids := tipSet.Get("Cids").Array()
		parentTipset := block["parentTipset"].(string)

		orphan := false
		for _, cidJSON := range cids {
			//log.Logger.Info("cid", "index", cidIndex, "val", gjson.Get(cidJSON.Raw, "/").String())
			id := gjson.Get(cidJSON.Raw, "/").String()

			//log.Logger.Info(id)
			if !strings.Contains(parentTipset, id) {
				orphan = true
				break
			}
		}

		if orphan {
			msg := client.Message{Type: client.OrphanBlock, Data: block}
			client.SendMessage(msg)
		}

		deleteKeys = append(deleteKeys, cid)

	}

	for _, v := range deleteKeys {
		delete(blockQueue, v)
	}
}
