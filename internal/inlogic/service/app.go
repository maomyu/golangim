package service

import (
	"context"

	"github.com/yuwe1/golangim/internal/inlogic/dao"
	"github.com/yuwe1/golangim/internal/model"
)

type appService struct {
}

var AppService = new(appService)

// 根据appId从数据库中获取app
func (a *appService) Get(ctx context.Context, appId int64) (*model.App, error) {
	// 先从缓存中获取，

	// 从数据库中获取
	app, err := dao.AppDao.GetApp(appId)
	if err != nil {
		return nil, err
	}

	//设置缓存
	if app != nil {

	}
	return app, nil

}
