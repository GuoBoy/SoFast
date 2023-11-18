package my_cmd

import (
	"SoFast/my_cmd/cmds"
	_ "SoFast/my_cmd/executer"
	"fmt"
	"strconv"
)

var (
	selKey string // 选项
)

func Execute() {
	for {
		fmt.Println("\n选择要执行的操作：")
		// 选项
		for _, option := range cmds.Cmds {
			fmt.Printf("\t【%d】 %s\n", option.ID, option.Description)
		}
		fmt.Print(">>>")
		if _, err := fmt.Scanln(&selKey); err != nil {
			continue
		}
		noUseSelKey := true
		for _, option := range cmds.Cmds {
			if selKey == strconv.Itoa(option.ID) {
				option.Execute()
				noUseSelKey = false
				fmt.Printf("\n*************Finished %s**************", option.Description)
				break
			}
		}
		if noUseSelKey {
			fmt.Print("选项错误!")
			continue
		}
	}
}
