package lib

import (
	"coffee-monitor/lib/client"
	fil "coffee-monitor/lib/fil/miner"
	"coffee-monitor/lib/shell"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

// Snapshot 快照一次
func Snapshot() {
	//连接服务端
	//go client.ConnectServer()
	//1.监控lotus数据并发送至服务端
	//err := LotusInfoCron()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//1.檢測lotus同步狀態
	go lotusCheck()

	//2.檢測lotus-miner info 並發送給server
	go lotusMinerInfoCheck()

	//3.监控miner日志发送块、孤块信息
	err := fil.MinerLogTailProcessor()
	if err != nil {
		fmt.Println(err)
	}

	//4.监控miner孤块信息
	go OrphanCheck()
}

// Snapshot 快照一次
func Lotus() {
	//1.檢測lotus同步狀態
	go lotusCheck()
	select {}
}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

//func LotusInfoCron() error {
//	log.Printf("start LotusInfoCron 0 */1 * * * ?")
//	c := newWithSeconds()
//	spec := "0 */1 * * * ?" //一分钟运行一次
//	_, err := c.AddFunc(spec, func() {
//		err := fil.SendLotusInfo()
//		if err != nil {
//			return
//		}
//	})
//	if err != nil {
//		return err
//	}
//	c.Start()
//	select {}
//
//}

func lotusCheck() {
	err := LotusSyncCron()
	if err != nil {
		log.Println(err)
		return
	}
}

func lotusMinerInfoCheck() {
	err := LotusMinerInfoCron()
	if err != nil {
		log.Println(err)
		return
	}
}

func OrphanCheck() {
	err := OrphanCheckCron()
	if err != nil {
		log.Println(err)
		return
	}
}

func OrphanCheckCron() error {
	log.Printf("start LotusSyncCron 0 */3 * * * ?")
	c := newWithSeconds()
	spec := "0 */3 * * * ?" //一分钟运行一次
	_, err := c.AddFunc(spec, func() {
		fil.CheckOrphanBlock()
	})
	if err != nil {
		return err
	}
	c.Start()
	select {}

}

func LotusSyncCron() error {
	log.Printf("@every 15s")
	c := newWithSeconds()
	//spec := "0 */1 * * * ?" //一分钟运行一次
	spec := "@every 15s"
	_, err := c.AddFunc(spec, func() {
		err := shell.LotusSyncCheck()
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}
	c.Start()
	select {}

}

func LotusMinerInfoCron() error {
	log.Printf("start LotusMinerInfoCron 0 */2 * * * ?")
	c := newWithSeconds()
	spec := "0 */2 * * * ?" //一分钟运行一次
	_, err := c.AddFunc(spec, func() {
		err, result := shell.LotusMinerInfo()
		if err != nil {
			log.Println(err)
			return
		}
		msg := client.Message{Type: client.LotusMinerInfo, Content: result}

		//time.Sleep(2000 * time.Millisecond)
		err2 := client.SendMessage(msg)
		if err2 != nil {
			log.Println(err)
			return
		}

	})
	if err != nil {
		return err
	}
	c.Start()
	select {}

}
