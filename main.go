package main

import (
	"coffee-monitor/lib"
	"fmt"
	"log"
	"strings"
	"time"
)

// 全局变量
var (
	bug_num int = 0 // 当日总bug数
)

// 主程序
func main() {
	diffTime := time.Now().UnixNano()
	defer func() {
		// 显示程序执行效率
		diffTime = (time.Now().UnixNano() - diffTime) / 1e6
		fmt.Printf("程序共执行 %v ms \n", diffTime)
	}()

	lib.FilInit()
	// 读配置文件
	config := lib.GetConfig()
	//fmt.Printf("%v\n\n", config)

	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Printf("处理的日期为：%v \n", date)

	for _, fileLogs := range config.Logfile {
		//api := fileLogs.Api

		//file := "/Users/fangyu/GolandProjects/coffee-monitor/logs/miner.log"
		blocks, err := lib.ReadNewBlocksFromLog(fileLogs.Path)
		if err != nil {
			log.Printf("err:%v", err)
			return
		}
		log.Printf("start filter invalid block...")
		forkedNum := 0
		for _, block := range blocks {
			_, err := lib.GetBlock(block["cid"].(string))
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
	// 发送邮件,耗时2秒多
	//libs.SendToMail(config.Mailto, "<h1>"+date+" BUG数汇总</h1><div>今日总bug数有"+strconv.Itoa(bug_num)+"个，请在 http://bugs.xxxxx.com/list?date="+date+" 中查看。</div>")

}
