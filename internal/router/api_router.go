package router

import (
	"github.com/gin-gonic/gin"

	"chat/internal/ctrl"
)

type Router struct {
	dingtalk *ctrl.DingtalkCtrl
}

func NewRouter(dingtalk *ctrl.DingtalkCtrl) *Router {
	return &Router{
		dingtalk: dingtalk,
	}
}

func (api *Router) RegisterChatRouter(r *gin.RouterGroup) {
	r.Any("/login", api.dingtalk.Login)
}
