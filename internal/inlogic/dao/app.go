package dao

import "github.com/yuwe1/golangim/internal/model"

import "github.com/yuwe1/golangim/basic/client/dbpool"

import "github.com/yuwe1/golangim/pkg/logger"

import "database/sql"

import "fmt"

type appDao struct {
}

var AppDao = new(appDao)

func (a *appDao) GetApp(appId int64) (*model.App, error) {
	s, err, p, c := dbpool.GetSession()
	defer func() {
		if s != nil {
			s.Relase(p, c)
		}
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	app := new(model.App)
	query := "select id,name,private_key,create_time,update_time from app where id = ?"
	s.Begin()
	err = s.QueryRow(query, appId).Scan(
		&app.Id, &app.Name, &app.PrivateKey,
		&app.CreateTime, &app.UpdateTime,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("%w", err)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	s.Commit()
	return app, nil
}
