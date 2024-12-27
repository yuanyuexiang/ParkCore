package controllers

import (
	"errors"
	"park/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	BaseController
}

// Logcat Logcat
// @Title Logcat
// @Summary 日志
// @Description 通过websocket推送日志
// @Success 200 {string}
// @router /logcat [get]
func (c *WebSocketController) Logcat() {
	log.Println("收到来自前端的websocket连接......")
	// Upgrade from http request to WebSocket.
	client, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		c.DoReturn(nil, errors.New("错误的连接"))
		return
	} else if err != nil {
		c.DoReturn(nil, err)
		return
	}
	key := uuid.NewV4().String()
	models.WebsocketClientMapForLogcat[key] = client
	defer client.Close()
	defer func() {
		delete(models.WebsocketClientMapForLogcat, key)
	}()
	log.Println("连接成功")
	for {
		_, message, err := client.ReadMessage()
		if err != nil {
			log.Println("读取数据错误:", err)
			break
		}
		log.Printf("读取数据: %s", message)
	}
}

// Information Information
// @Title Information
// @Summary 数据
// @Description 通过websocket推送数据
// @Success 200 {object} models.Information
// @router /information [get]
func (c *WebSocketController) Information() {
	log.Println("收到来自前端websocket的连接......")
	// Upgrade from http request to WebSocket.
	client, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		c.DoReturn(nil, errors.New("错误的连接"))
		return
	} else if err != nil {
		c.DoReturn(nil, err)
		return
	}
	key := uuid.NewV4().String()
	models.WebsocketClientMapForData[key] = client
	defer client.Close()
	defer func() {
		delete(models.WebsocketClientMapForData, key)
	}()
	log.Println("连接成功")
	for {
		_, message, err := client.ReadMessage()
		if err != nil {
			log.Println("读取数据错误:", err)
			break
		}
		log.Printf("收到数据: %s", message)
	}
}
