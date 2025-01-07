package models

import (
	"encoding/json"
	"log"
	"park/utils"

	"github.com/gorilla/websocket"
)

var (
	WebsocketClientMapForData   = make(map[string]*websocket.Conn)
	WebsocketClientMapForLogcat = make(map[string]*websocket.Conn)
	InformationChan             = make(chan *Information, 500) //发向前端的消息队列
	LogChan                     = make(chan string, 500)       //发向前端的消息队列
)

// Information Information
type Information struct {
	Type    int         `json:"type" description:"数据类型"`
	Message string      `json:"message" description:"信息"`
	Data    interface{} `json:"data,omitempty"`
}

func init() {
	go messageHandle()
}

// 处理引擎
func messageHandle() {
	utils.Println("启动消息处理引擎......")
	for {
		select {
		case message := <-InformationChan:
			//utils.PrintlnWithTag("推送数据", message)
			for _, websocketClient := range WebsocketClientMapForData {
				b, _ := json.Marshal(message)
				err := websocketClient.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					log.Println("向前端写数据出错:", err)
					break
				}
			}
		case message := <-LogChan:
			utils.PrintlnWithTag("推送日志", message)
			for _, websocketClient := range WebsocketClientMapForLogcat {
				b, _ := json.Marshal(message)
				err := websocketClient.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					log.Println("向前端写数据出错:", err)
					break
				}
			}
		}
	}
}

// 向订阅者发送日志
func PushLog(log string) {
	LogChan <- log
}

// 向订阅者发送日志
func PushLogX(log interface{}) {
	o, _ := json.MarshalIndent(log, "", "    ")
	LogChan <- string(o)
}

// 向订阅者发送数据
func PushData(information *Information) {
	InformationChan <- information
}
