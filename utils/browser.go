package utils

import (
	"fmt"
	"os/exec"
)

// OpenUrl 打开浏览器
func OpenUrl(addr string) {
	err := exec.Command("cmd.exe", "/c", "start "+addr).Start()
	if err != nil {
		fmt.Println(err)
	}
}
