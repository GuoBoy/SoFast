package server

import (
	"SoFast/config"
	"SoFast/handles"
	"SoFast/middleware"
	"SoFast/static"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
)

func Run() {
	c := config.Config
	var r *gin.Engine
	if c.Dev {
		r = gin.Default()
		r.Use(middleware.Cors())
		r.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", nil)
		})
	} else {
		r = gin.New()
		r.Use(gin.Recovery())
		r.Any("/", func(context *gin.Context) {
			context.Writer.Write(static.IndexFile)
		})
		assets, _ := fs.Sub(static.Assets, "dist/assets")
		r.StaticFS("/assets", http.FS(assets))
	}
	r.GET("/ws", handles.WsHandle)
	r.GET("/files", handles.GetFiles)
	r.POST("/upload", handles.UploadHandle)
	r.DELETE("/delete/:filename", handles.DeleteFile)
	r.StaticFS("/download", http.Dir(c.UploadPath))
	r.StaticFile("/addr", "addr.png")
	log.Fatal(r.Run(fmt.Sprintf("0.0.0.0:%d", c.Port)))
}
