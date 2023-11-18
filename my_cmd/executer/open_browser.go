package executer

import (
	"SoFast/config"
	"SoFast/my_cmd/cmds"
	"SoFast/utils"
)

func Execute() {
	addr := config.GetServerAddr()
	utils.OpenUrl(addr.Addr)
}

func init() {
	cmds.InstallCmd(cmds.Cmd{ID: 1, Description: "打开浏览器", Execute: Execute})
}
