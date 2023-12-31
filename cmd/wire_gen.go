// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"chat/configs"
	"chat/internal/ctrl"
	"chat/internal/router"
	"chat/internal/server"
	"chat/internal/service"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

// initApplication init application.
func initApplication(leve string, app *configs.Applet) (*gin.Engine, error) {
	dingtalkService, err := service.NewDingtalkService(app)
	if err != nil {
		return nil, err
	}
	dingtalkCtrl := ctrl.NewDingtalkCtrl(dingtalkService)
	routerRouter := router.NewRouter(dingtalkCtrl)
	engine := server.NewHttpServer(leve, routerRouter)
	return engine, nil
}
