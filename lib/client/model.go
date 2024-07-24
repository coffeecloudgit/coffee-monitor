package client

type MsgType string

const (
	NewBlock       MsgType = "new-block"
	OrphanBlock    MsgType = "orphan-block"
	WARNING        MsgType = "warning"
	LotusMinerInfo MsgType = "lotus-miner-info"
)

type Message struct {
	Type    MsgType                `json:"type"`    //消息类型
	Content string                 `json:"content"` //消息内容
	Data    map[string]interface{} `json:"data"`    //{type:'new_block', data:{}} //消息包含数据
}
