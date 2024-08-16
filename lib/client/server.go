package client

import (
	"coffee-monitor/lib/log"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var serverAddr = flag.String("server", "0.0.0.0:8083", "http service address")
var ClientsChecks = make(map[string]int64, 2)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func IsOnline(miner string) bool {
	if val, ok := ClientsChecks[miner]; ok {
		if (time.Now().UnixMilli() - val) < 60000 {
			return true
		}
	}

	return false
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Info("upgrade:", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Logger.Info("read:", err)
			break
		}

		stringMsg := string(message)
		if !strings.HasPrefix(strings.TrimSpace(stringMsg), "{") {
			log.Logger.Error("error msg", "msg", stringMsg)
			continue
		}

		//log.Logger.Info("recv: aaa|%s|bb, type: %v \n", message, mt)
		//{"type":"new-mine-one","content":"","data":{"epoch":3888596,"miner":"f02246008"}}
		var msg Message
		err2 := json.Unmarshal(message, &msg)
		log.Logger.Info("msg:", slog.String("data", string(message)))
		if err2 != nil {
			log.Logger.Info(err2.Error())
			continue
		}

		if msg.Type == Ping {
			pongMsg := fmt.Sprintf("{\"type\":\"%s\", \"content\": \"%s\"}", Pong, msg.Content)
			log.Logger.Debug("connected check ...")
			err = c.WriteMessage(mt, []byte(pongMsg))
			if err != nil {
				log.Logger.Error("write:", "err:", err)
				continue
			}

			ClientsChecks[msg.Content] = time.Now().UnixMilli()
		} else {
			processMsg(&msg)
		}
	}
}

// {"type":"lotus-miner-info","content":"Enabled subsystems: [Mining Sealing SectorStorage]\nStartTime: 314h41m30s (started at 2024-07-10 07:52:23 +0800 CST)\nChain: [sync ok] [basefee 100 aFIL]\n⚠ 1 Active alerts (check lotus-miner log alerts)\nMiner: f03148950 (32 GiB sectors)\nPower: 9.91 Pi / 22.7 Ei (0.0427%)\n        Raw: 1015 TiB / 5.662 EiB (0.0171%)\n        Committed: 1.063 PiB\n        Proving: 1015 TiB\nProjected average block win rate: 42.98/week (every 3h54m31s)\nProjected block win with 99.9% probability every 26h58m17s\n(projections DO NOT account for future network and miner growth)\n\nMiner Balance:    55398.537 FIL\n      PreCommit:  7.056 FIL\n      Pledge:     52371.423 FIL\n      Vesting:    155.325 FIL\n      Available:  2864.733 FIL\nMarket Balance:   255 FIL\n       Locked:    231.44 FIL\n       Available: 23.56 FIL\nWorker Balance:   34.359 FIL\n       Control:   88.59 FIL\nTotal Spendable:  3011.242 FIL\n\nBeneficiary:    f03155156\n\nSectors:\n        Total: 35660\n        Proving: 34898\n        AddPiece: 14\n        PreCommit1: 474\n        PreCommit2: 70\n        PreCommitBatchWait: 6\n        WaitSeed: 169\n        Committing: 11\n        CommitAggregateWait: 9\n        FinalizeSector: 7\n        Removed: 2\n\nWorkers: Seal(77) WdPoSt(1) WinPoSt(0)","data":null},

func processMsg(msg *Message) {
	switch msg.Type {
	case NewBlock:
	case OrphanBlock:
	case LotusMinerInfo:
	case NewMineOne:

	default:
		fmt.Println("未知消息：", msg)
	}
}

//func home(w http.ResponseWriter, r *http.Request) {
//	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
//}

func RunTestServer() {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	//http.HandleFunc("/", home)
	http.ListenAndServe(*serverAddr, nil)
}

//
//var homeTemplate = template.Must(template.New("").Parse(`
//<!DOCTYPE html>
//<html>
//<head>
//<meta charset="utf-8">
//<script>
//window.addEventListener("load", function(evt) {
//
//    var output = document.getElementById("output");
//    var input = document.getElementById("input");
//    var ws;
//
//    var print = function(message) {
//        var d = document.createElement("div");
//        d.textContent = message;
//        output.appendChild(d);
//        output.scroll(0, output.scrollHeight);
//    };
//
//    document.getElementById("open").onclick = function(evt) {
//        if (ws) {
//            return false;
//        }
//        ws = new WebSocket("{{.}}");
//        ws.onopen = function(evt) {
//            print("OPEN");
//        }
//        ws.onclose = function(evt) {
//            print("CLOSE");
//            ws = null;
//        }
//        ws.onmessage = function(evt) {
//            print("RESPONSE: " + evt.data);
//        }
//        ws.onerror = function(evt) {
//            print("ERROR: " + evt.data);
//        }
//        return false;
//    };
//
//    document.getElementById("send").onclick = function(evt) {
//        if (!ws) {
//            return false;
//        }
//        print("SEND: " + input.value);
//        ws.send(input.value);
//        return false;
//    };
//
//    document.getElementById("close").onclick = function(evt) {
//        if (!ws) {
//            return false;
//        }
//        ws.close();
//        return false;
//    };
//
//});
//</script>
//</head>
//<body>
//<table>
//<tr><td valign="top" width="50%">
//<p>Click "Open" to create a connection to the server,
//"Send" to send a message to the server and "Close" to close the connection.
//You can change the message and send multiple times.
//<p>
//<form>
//<button id="open">Open</button>
//<button id="close">Close</button>
//<p><input id="input" type="text" value="Hello world!">
//<button id="send">Send</button>
//</form>
//</td><td valign="top" width="50%">
//<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
//</td></tr></table>
//</body>
//</html>
//`))
