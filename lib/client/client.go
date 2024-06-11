package client

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"sync"
	"time"
)

var connectServerAddr = flag.String("addr", "ws://127.0.0.1:8083/echo", "http service address")
var wsClient *Wsc = nil
var connectLck sync.Mutex

func DisConnectServer() {
	if wsClient == nil {
		return
	}
	wsClient.Close()
}
func ConnectServer() {
	connectLck.Lock()
	defer connectLck.Unlock()

	if wsClient != nil {
		return
	}

	flag.Parse()
	log.SetFlags(0)

	done := make(chan bool)
	wsClient = NewWsClient(*connectServerAddr)
	// 可自定义配置，不使用默认配置
	//wsClient.SetConfig(&wsc.Config{
	//	// 写超时
	//	WriteWait: 10 * time.Second,
	//	// 支持接受的消息最大长度，默认512字节
	//	MaxMessageSize: 2048,
	//	// 最小重连时间间隔
	//	MinRecTime: 2 * time.Second,
	//	// 最大重连时间间隔
	//	MaxRecTime: 60 * time.Second,
	//	// 每次重连失败继续重连的时间间隔递增的乘数因子，递增到最大重连时间间隔为止
	//	RecFactor: 1.5,
	//	// 消息发送缓冲池大小，默认256
	//	MessageBufferSize: 1024,
	//})
	// 设置回调处理
	wsClient.OnConnected(func() {
		log.Println("OnConnected: ", wsClient.WebSocket.Url)
		// 连接成功后，测试每10秒发送消息
		go func() {
			t := time.NewTicker(10 * time.Second)
			for {
				select {
				case <-t.C:
					err := wsClient.SendTextMessage("hello")
					if err == CloseErr {
						return
					}
				}
			}
		}()
	})
	wsClient.OnConnectError(func(err error) {
		log.Println("OnConnectError: ", err.Error())
	})
	wsClient.OnDisconnected(func(err error) {
		log.Println("OnDisconnected: ", err.Error())
	})
	wsClient.OnClose(func(code int, text string) {
		log.Println("OnClose: ", code, text)
		done <- true
	})
	//wsClient.OnTextMessageSent(func(message string) {
	//	//log.Println("OnTextMessageSent: ", message)
	//})
	wsClient.OnBinaryMessageSent(func(data []byte) {
		log.Println("OnBinaryMessageSent: ", string(data))
	})
	wsClient.OnSentError(func(err error) {
		log.Println("OnSentError: ", err.Error())
	})
	wsClient.OnPingReceived(func(appData string) {
		log.Println("OnPingReceived: ", appData)
	})
	wsClient.OnPongReceived(func(appData string) {
		log.Println("OnPongReceived: ", appData)
	})
	wsClient.OnTextMessageReceived(func(message string) {
		log.Println("OnTextMessageReceived: ", message)
	})
	wsClient.OnBinaryMessageReceived(func(data []byte) {
		log.Println("OnBinaryMessageReceived: ", string(data))
	})
	// 开始连接
	go wsClient.Connect()
	for {
		select {
		case <-done:
			return
		}
	}

}

func SendData(message string) error {
	if wsClient == nil {
		ConnectServer()
	}

	return wsClient.SendTextMessage(message)
}

func SendMessage(message Message) error {
	if wsClient == nil {
		ConnectServer()
	}
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return errors.New("message marshal fail")
	}
	log.Println(string(msgBytes))
	return wsClient.SendTextMessage(string(msgBytes))
}
