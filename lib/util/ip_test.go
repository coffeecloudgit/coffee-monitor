package util

import (
	"log"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	got := GetLocalIP()

	log.Println(got)
}
