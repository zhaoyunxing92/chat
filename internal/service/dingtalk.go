package service

import (
	"github.com/zhaoyunxing92/dingtalk/v2"
	"github.com/zhaoyunxing92/dingtalk/v2/response"

	"chat/configs"
	"chat/internal/domain"
	"chat/pkg/log"
)

type DingtalkService struct {
	client *dingtalk.DingTalk
}

func NewDingtalkService(app *configs.Applet) (*DingtalkService, error) {
	client, err := dingtalk.NewClient(app.Key, app.Secret)
	if err != nil {
		log.Errorf("dingtalk.NewClient err:%v", err)
		return nil, err
	}
	return &DingtalkService{client: client}, nil
}

func (svc *DingtalkService) Login(code string) (*domain.Employee, error) {
	var (
		err error
		res response.CodeGetUserInfo
	)

	if res, err = svc.client.GetUserInfoByCode(code); err != nil {

	}

	return &domain.Employee{
		Name:     res.UserInfo.Name,
		UnionId:  res.UserInfo.UnionId,
		UserId:   res.UserInfo.UserId,
		Admin:    res.UserInfo.Admin,
		DeviceId: res.UserInfo.DeviceId,
	}, nil
}
