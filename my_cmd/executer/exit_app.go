package executer

import (
	"SoFast/my_cmd/cmds"
	"os"
)

func init() {
	cmds.InstallCmd(cmds.Cmd{
		ID:          2,
		Description: "退出系统",
		Execute: func() {
			os.Exit(0)
		},
	})
}
