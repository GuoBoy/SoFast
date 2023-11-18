package utils

import (
	"log"
	"net"
	"strings"
)

var ip string

// GetLocalIP 获取本机IP
func GetLocalIP() string {
	if ip != "" {
		return ip
	}
	//host, _ := os.Hostname()
	//fmt.Println(host)
	//ip, err := net.LookupIP(host)
	//if err != nil {
	//	fmt.Println(err)
	//	return ""
	//}
	//for _, i := range ip {
	//	fmt.Println(i.String())
	//}
	//return ip[len(ip)-1].String()
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ip = strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip
}
