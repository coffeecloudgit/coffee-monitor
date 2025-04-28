package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func Test_notice(t *testing.T) {
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/2 * * * * ?"

	c.Start()
	defer c.Stop()

	_, err := c.AddFunc(spec, func() {
		i++
		fmt.Println("cron times : ", i)
	})

	if err != nil {
		fmt.Errorf("AddFunc error : %v", err)
		return
	}
	select {}
}

func Test_cron_parse(t *testing.T) {
	spec := "36 14 * * *" // 如果配置为空，默认每天3点运行一次

	fmt.Println("start SectorsExpireInfoCron " + spec)
	c := cron.New()
	_, err2 := c.AddFunc(spec, func() {
		fmt.Println("cron times : ", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err2 != nil {
		fmt.Printf("AddFunc error : %v \n", err2)
		//return
	}
	c.Start()
	defer c.Stop()
	fmt.Println("cron add success")
	time.Sleep(200 * time.Second)
	//select {}

}
