package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chat/internal/service"
	"chat/pkg/log"
)

type DingtalkCtrl struct {
	service *service.DingtalkService
}

func NewDingtalkCtrl(service *service.DingtalkService) *DingtalkCtrl {
	return &DingtalkCtrl{
		service: service,
	}
}

func (ctrl *DingtalkCtrl) Login(ctx *gin.Context) {
	code := ctx.GetString("code")
	log.Infof("dingtalk code: %s", code)
	if employee, err := ctrl.service.Login(code); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": employee,
		})
	}
}
