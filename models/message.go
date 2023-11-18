package models

import (
	"time"
)

const (
	Offline       = iota // 下线
	Online               // 上线
	SuccessOnline        // 上线成功
	Text                 // 文本消息
	MsgOk                // 消息发送成功
	Other                // 未知类型，作为字符串解析
	UserList             // 获取用户列表
)

type Message struct {
	Username   string      `json:"username"`
	MsgType    int         `json:"msg_type"`
	Data       interface{} `json:"data"`
	CreateTime string      `json:"created_at"`
}

func NewMessage(username string, msgType int, data interface{}) Message {
	return Message{
		Username:   username,
		MsgType:    msgType,
		Data:       data,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
}
