package client

type MsgType string

const (
	Ping           MsgType = "ping"
	Pong           MsgType = "pong"
	NewMineOne     MsgType = "new-mine-one"
	NewBlock       MsgType = "new-block"
	OrphanBlock    MsgType = "orphan-block"
	LotusMinerInfo MsgType = "lotus-miner-info"
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
