package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	RunTestServer()
}

func TestJson2(t *testing.T) {
	message := "{\"type\":\"new-mine-one\",\"content\":\"\",\"data\":{\"epoch\":3888596,\"miner\":\"f02246008\"}}"
	var msg2 Message
	err := json.Unmarshal([]byte(message), &msg2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg2)
}
