package server

import (
	"github.com/gin-gonic/gin"

	"chat/internal/middleware"
	"chat/internal/router"
)

func NewHttpServer(level string, router *router.Router) *gin.Engine {
	if level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.GET("/health", func(ctx *gin.Context) { ctx.String(200, "OK") })

	chat := r.Group("/chat/api/v1")
	router.RegisterChatRouter(chat)
	return r
}
