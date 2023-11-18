package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	Conn       *websocket.Conn
	RemoteAddr string `json:"remote_addr"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
}
