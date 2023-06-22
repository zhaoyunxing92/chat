package internal

import (
	"github.com/google/wire"

	"chat/internal/ctrl"
	"chat/internal/router"
	"chat/internal/server"
	"chat/internal/service"
)

var Provider = wire.NewSet(
	router.NewRouter,
	ctrl.NewDingtalkCtrl,
	service.NewDingtalkService,
	server.NewHttpServer,
)
