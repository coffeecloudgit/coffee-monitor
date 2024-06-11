package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
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
