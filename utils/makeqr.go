package utils

import (
	"github.com/skip2/go-qrcode"
)

// MakeQr 生成二维码并开启网页
func MakeQr(addr string) {
	err := qrcode.WriteFile(addr, qrcode.High, 1024, "./addr.png")
	if err != nil {
		panic(err)
	}
	// 展示二维码图片
	//err = exec.Command("cmd.exe", "/c", "start addr.png").Start()
	//if err != nil {
	//	panic(err)
	//}
}
