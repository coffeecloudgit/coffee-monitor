package main

import (
	"coffee-monitor/lib/cmd"
	"coffee-monitor/lib/log"
)

var Version = "unknown"

// 主程序
func main() {
	log.Logger.Info("Program info", "version", Version)
	//启动命令行
	cmd.Execute()
	// 发送邮件,耗时2秒多
	//libs.SendToMail(config.Mailto, "<h1>"+date+" BUG数汇总</h1><div>今日总bug数有"+strconv.Itoa(bug_num)+"个，请在 http://bugs.xxxxx.com/list?date="+date+" 中查看。</div>")
	log.Logger.Info("end")
}
