package handles

import (
	"SoFast/config"
	"github.com/gin-gonic/gin"
	"path"
)

func UploadHandle(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(402, gin.H{
			"code": 402,
			"msg":  "上传失败" + err.Error(),
		})
		return
	}
	err = ctx.SaveUploadedFile(file, path.Join(config.Config.UploadPath, file.Filename))
	if err != nil {
		ctx.JSON(402, gin.H{
			"code": 402,
			"msg":  "上传失败" + err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
