package util

import (
	"coffee-monitor/lib/log"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	got := GetLocalIP()

	log.Logger.Info(got)
}
