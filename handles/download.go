package handles

import (
	"SoFast/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func GetFiles(ctx *gin.Context) {
	files, err := os.ReadDir(config.Config.UploadPath)
	if err != nil {
		fmt.Println("读取错误", err)
		ctx.JSON(402, gin.H{
			"msg": err.Error(),
		})
	}
	var res []map[string]interface{}
	// TODO return file tree
	for _, file := range files {
		res = append(res, map[string]interface{}{"filename": file.Name(), "filesize": file})
	}
	ctx.JSON(200, gin.H{
		"code":   200,
		"length": len(res),
		"data":   res,
	})
}
