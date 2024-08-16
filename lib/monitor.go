package lib

import (
	"coffee-monitor/lib/client"
	config2 "coffee-monitor/lib/config"
	fil "coffee-monitor/lib/fil/miner"
	"coffee-monitor/lib/log"
	"coffee-monitor/lib/shell"
	"github.com/robfig/cron/v3"
)

// Snapshot 快照一次
func Snapshot() {
	//连接服务端
	//go client.ConnectServer()
	//1.监控lotus数据并发送至服务端
	//err := LotusInfoCron()
	//if err != nil {
	//	log.Logger.Fatal(err)
	//	return
	//}
	//1.檢測lotus同步狀態
	go lotusCheck()

	//2.檢測lotus-miner info 並發送給server
	go lotusMinerInfoCheck()

	//3.监控miner日志发送块、孤块信息
	go minerLogCheck()
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
//	log.Logger.Info("start LotusInfoCron 0 */1 * * * ?")
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
		log.Logger.Info(err.Error())
		return
	}
}

func lotusMinerInfoCheck() {
	err := LotusMinerInfoCron()
	if err != nil {
		log.Logger.Info(err.Error())
		return
	}
}

func minerLogCheck() {
	err := fil.MinerLogTailProcessor()
	if err != nil {
		log.Logger.Info(err.Error())
	}
}

func OrphanCheck() {
	err := OrphanCheckCron()
	if err != nil {
		log.Logger.Info(err.Error())
		return
	}
}

func OrphanCheckCron() error {
	log.Logger.Info("start OrphanCheckCron 0 */1 * * * ?")
	c := newWithSeconds()
	spec := "0 */1 * * * ?" //一分钟运行一次
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
	log.Logger.Info("@every 20s")
	c := newWithSeconds()
	//spec := "0 */1 * * * ?" //一分钟运行一次
	spec := "@every 20s"
	_, err := c.AddFunc(spec, func() {
		err := shell.LotusSyncCheck()
		if err != nil {
			log.Logger.Info(err.Error())
		}
	})
	if err != nil {
		return err
	}
	c.Start()
	select {}

}

func LotusMinerInfoCron() error {
	log.Logger.Info("start LotusMinerInfoCron 0 */2 * * * ?")
	c := newWithSeconds()
	spec := "0 */2 * * * ?" //一分钟运行一次
	_, err := c.AddFunc(spec, func() {
		err, result := shell.LotusMinerInfo()
		if err != nil {
			log.Logger.Info(err.Error())
			return
		}
		mine := make(map[string]interface{}, 5)
		mine["miner"] = config2.CONF.Fil.Account

		msg := client.Message{Type: client.LotusMinerInfo, Content: result, Data: mine}

		//time.Sleep(2000 * time.Millisecond)
		err2 := client.SendMessage(msg)
		if err2 != nil {
			log.Logger.Info(err2.Error())
			return
		}

	})
	if err != nil {
		return err
	}
	c.Start()
	select {}

}
