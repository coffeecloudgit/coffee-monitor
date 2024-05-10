package fil

// LotusInfo lotus 监测数据
type LotusInfo struct {
	Id       string `json:"id"`       //node ID
	Height   uint64 `json:"height"`   //同步高度
	PeersNum int    `json:"peersNum"` //连接的节点数量
	Account  string `json:"account"`  //关联账户，如：f02246008
	Ip       string `json:"ip"`       //主机IP 如：192.168.0.2
}

type Miner struct {
}

type ForkedBlockInfo struct {
	totalNum   uint64
	forkedNum  uint64
	forkedRate float64
}
