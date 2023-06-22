//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"chat/configs"
	"chat/internal"
)

// initApplication init application.
func initApplication(leve string, app *configs.Applet) (*gin.Engine, error) {
	panic(wire.Build(internal.Provider))
}
