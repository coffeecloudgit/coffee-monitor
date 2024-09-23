package client

import "fmt"

type MsgType string

const (
	Ping              MsgType = "ping"
	Pong              MsgType = "pong"
	NewMineOne        MsgType = "new-mine-one"
	NewBlock          MsgType = "new-block"
	OrphanBlock       MsgType = "orphan-block"
	LotusMinerInfo    MsgType = "lotus-miner-info"
	SectorsExpireInfo MsgType = "sectors-expire-info"
)

var (
	PingJson = "{\"type\":\"%s\", \"content\": \"%s\", \"version\": \"%s\"}"
	//PingJson = fmt.Sprintf("{\"type\":\"%s\", \"content\": \"%s\", \"version\": \"%s\"}", Ping, config2.CONF.Fil.Account, &program.Version)
)

type Message struct {
	Type    MsgType                `json:"type"`    //消息类型
	Content string                 `json:"content"` //消息内容
	Data    map[string]interface{} `json:"data"`    //{type:'new_block', data:{}} //消息包含数据
}

// ExpireSameDaySectors 同一天到期的sector
type ExpireSameDaySectors struct {
	Day       string `json:"day"`  //如：2023-01-01
	From      uint64 `json:"from"` //
	To        uint64 `json:"to"`   //
	SectorNum uint64 `json:"sectorNum"`
}

func (f *ExpireSameDaySectors) String() string {
	return fmt.Sprintf("{day: %s, from:%d, to:%d, sectorNum:%d}", f.Day, f.From, f.To, f.SectorNum)
}
