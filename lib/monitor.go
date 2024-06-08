package lib

import (
	"coffee-monitor/lib/fil"
	"coffee-monitor/lib/shell"
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

	err := LotusSyncCron()
	if err != nil {
		log.Println(err)
		return
	}

	//2.监控miner日志发送孤块信息

	//3.监控miner数据并发送至服务端

}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func LotusInfoCron() error {
	log.Printf("start LotusInfoCron 0 */1 * * * ?")
	c := newWithSeconds()
	spec := "0 */1 * * * ?" //一分钟运行一次
	_, err := c.AddFunc(spec, func() {
		err := fil.SendLotusInfo()
		if err != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	c.Start()
	select {}

}

func LotusSyncCron() error {
	log.Printf("start LotusSyncCron 0 */1 * * * ?")
	c := newWithSeconds()
	spec := "0 */1 * * * ?" //一分钟运行一次
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
