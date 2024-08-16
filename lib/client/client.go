package client

import (
	"coffee-monitor/lib/config"
	"coffee-monitor/lib/log"
	"encoding/json"
	"errors"
	"flag"
	"log/slog"
	"strings"
	"sync"
	"time"
)

var connectServerAddr = flag.String("addr", config.CONF.Fil.MsgServer, "http service address")
var wsClient *Wsc = nil
var connectLck sync.Mutex

func DisConnectServer() {
	if wsClient == nil {
		return
	}
	wsClient.Close()
}

func IsConnectLocalhostServer() bool {
	if strings.Contains(*connectServerAddr, "127.0.0.1:8083") || strings.Contains(*connectServerAddr, "localhost:8083") {
		return true
	}
	return false
}
func ConnectServer() {
	connectLck.Lock()
	defer connectLck.Unlock()

	if wsClient != nil {
		return
	}

	flag.Parse()
	//log.Logger.SetFlags(0)

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
		log.Logger.Info("OnConnected: ", slog.String("url", wsClient.WebSocket.Url))
		// 连接成功后，测试每10秒发送消息
		go func() {
			t := time.NewTicker(20 * time.Second)
			for {
				select {
				case <-t.C:
					err := wsClient.SendTextMessage(PingJson)
					if err == CloseErr {
						return
					}
				}
			}
		}()
	})
	wsClient.OnConnectError(func(err error) {
		log.Logger.Info("OnConnectError: ", slog.String("error", err.Error()))
	})
	wsClient.OnDisconnected(func(err error) {
		log.Logger.Info("OnDisconnected: %s", slog.String("error", err.Error()))
	})
	wsClient.OnClose(func(code int, text string) {
		log.Logger.Info("OnClose: ", slog.Int("code", code), slog.String("text", text))
		done <- true
	})
	//wsClient.OnTextMessageSent(func(message string) {
	//	//log.Logger.Info("OnTextMessageSent: ", message)
	//})
	wsClient.OnBinaryMessageSent(func(data []byte) {
		log.Logger.Info("OnBinaryMessageSent: ", slog.String("data", string(data)))
	})
	wsClient.OnSentError(func(err error) {
		log.Logger.Info("OnSentError:", slog.String("error", err.Error()))
	})
	wsClient.OnPingReceived(func(appData string) {
		log.Logger.Info("OnPingReceived: ", slog.String("data", appData))
	})
	wsClient.OnPongReceived(func(appData string) {
		log.Logger.Info("OnPongReceived: ", slog.String("data", appData))
	})
	wsClient.OnTextMessageReceived(func(message string) {
		log.Logger.Debug("OnTextMessageReceived:", slog.String("message", message))
	})
	wsClient.OnBinaryMessageReceived(func(data []byte) {
		log.Logger.Info("OnBinaryMessageReceived: ", slog.String("data", string(data)))
	})
	// 开始连接
	wsClient.Connect()
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
	log.Logger.Info(string(msgBytes))
	return wsClient.SendTextMessage(string(msgBytes))
}
