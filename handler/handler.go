package handler

import (
	"github.com/Shinji-Kubo/sopa/server"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.LoadHTMLFiles("templates/index.html")

	r.GET("/", server.RootPath)

	return r
}
