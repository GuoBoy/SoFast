package main

import (
	"SoFast/config"
	"SoFast/hub"
	"SoFast/my_cmd"
	"SoFast/server"
	"SoFast/utils"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	c := config.Init()
	addr := config.GetServerAddr()
	fmt.Println("SoFast is listening at " + addr.Addr)
	// run
	go hub.RunBroadcast()
	go server.Run()
	if !c.Dev {
		utils.MakeQr(addr.Addr)
		utils.OpenUrl(addr.Addr)
	}
	my_cmd.Execute()
}
