package handles

import (
	"SoFast/hub"
	"SoFast/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:   1024 * 10,
	WriteBufferSize:  1024 * 10,
	HandshakeTimeout: time.Second * 10,
}

func WsHandle(ctx *gin.Context) {
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	cip := ctx.ClientIP()
	user := models.User{
		Conn:       conn,
		RemoteAddr: cip,
		Username:   cip,
		Avatar:     "",
	}
	hub.AddWsUser(user)
	for {
		var msg models.Message
		if err = conn.ReadJSON(&msg); err != nil {
			hub.RemoveUser(user.RemoteAddr)
			hub.BroadcastMessage(models.NewMessage(cip, models.Offline, nil))
			fmt.Println("解析请求失败: ", err)
			return
		}
		switch msg.MsgType {
		case models.Online:
			// 用户上线
			if err = conn.WriteJSON(models.NewMessage(cip, models.SuccessOnline, nil)); err != nil {
				fmt.Printf("反馈给%s上线成功消息失败:%t", user.Username, err)
				return
			}
			// 发送用户列表
			var nowUsers []string
			for _, u := range hub.AllUser() {
				nowUsers = append(nowUsers, fmt.Sprintf("%s^%s", u.Username, u.RemoteAddr))
			}
			hub.BroadcastMessage(models.NewMessage(cip, models.UserList, nowUsers))
			break
		case models.Offline:
			hub.BroadcastMessage(models.NewMessage(cip, models.Offline, nil))
			return
		case models.Text:
			hub.BroadcastMessage(msg)
			break
		case models.UserList:
			var nowUsers []string
			for _, u := range hub.AllUser() {
				nowUsers = append(nowUsers, fmt.Sprintf("%s^%s", u.Username, u.RemoteAddr))
			}
			if err = conn.WriteJSON(models.NewMessage(cip, models.UserList, nowUsers)); err != nil {
				fmt.Println("发送用户列表失败")
			}
			break
		default:
			continue
		}
	}
}
