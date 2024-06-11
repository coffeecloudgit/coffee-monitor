package client

type MsgType string

const (
	NEW_BLOCK    MsgType = "new_block"
	ORPHAN_BLOCK MsgType = "orphan_block"
	WARNING      MsgType = "warning"
)

type Message struct {
	Type    MsgType                `json:"type"`    //消息类型
	Content string                 `json:"content"` //消息内容
	Data    map[string]interface{} `json:"data"`    //{type:'new_block', data:{}} //消息包含数据
}
