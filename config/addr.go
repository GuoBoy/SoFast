package config

import (
	"SoFast/utils"
	"fmt"
)

type ServerAddr struct {
	Host string
	Port uint
	Addr string
}

func GetServerAddr() ServerAddr {
	host := utils.GetLocalIP()
	port := Config.Port
	return ServerAddr{
		Host: host,
		Port: port,
		Addr: fmt.Sprintf(host, port),
	}
}
