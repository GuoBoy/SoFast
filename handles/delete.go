package handles

import (
	"SoFast/config"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func DeleteFile(ctx *gin.Context) {
	name := ctx.Param("filename")
	if name == "" {
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  "参数错误",
		})
		return
	}
	filename := filepath.Join(config.Config.UploadPath, name)
	_, err := os.Stat(filename)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 404,
			"msg":  err.Error(),
		})
		return
	}
	if err = os.Remove(filename); err != nil {
		ctx.JSON(200, gin.H{
			"code": 403,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
	})
}
