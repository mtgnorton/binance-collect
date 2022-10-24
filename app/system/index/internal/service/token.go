package service

import (
	"context"
	"gf-admin/app/model"
	"gf-admin/app/shared"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var FrontTokenInstance = frontTokenHandle{ //如果直接在这里获取配置，会出错，因为这里的执行顺序可能早于boot.go加载配置文件，所以改为使用LoadConfig方法
}

type frontTokenHandle struct {
	shared.TokenHandler
	user       *model.UserSummary
	loadConfig bool
}

//加载配置文件里面token的相关配置,应该最先调用该方法
func (a *frontTokenHandle) LoadConfig() *frontTokenHandle {
	ctx := context.TODO()
	FrontTokenInstance.TokenHandler = shared.TokenHandler{
		CacheMode:  shared.CacheModeRedis,
		Timeout:    g.Cfg().MustGet(ctx, "front_token.Timeout").Int(),
		MaxRefresh: g.Cfg().MustGet(ctx, "front_token.MaxRefresh").Int(),
		CacheKey:   g.Cfg().MustGet(ctx, "front_token.CacheKey").String(),
		EncryptKey: g.Cfg().MustGet(ctx, "front_token.EncryptKey").Bytes(),
		MultiLogin: g.Cfg().MustGet(ctx, "front_token.MultiLogin").Bool(),
	}
	a.loadConfig = true
	return a
}

func (a *frontTokenHandle) GetUser(ctx context.Context) (user *model.UserSummary, err error) {

	if !a.loadConfig {
		a.LoadConfig()
	}

	if a.user != nil && a.user.Id != 0 {
		return a.user, nil
	}
	data := shared.Context.GetUser(ctx)
	g.Dump("getUser", data)
	user = &model.UserSummary{}
	err = gconv.Scan(data, &user)
	a.user = user
	return
}

func (a *frontTokenHandle) GetUserId(ctx context.Context) (userId uint, err error) {
	user, err := a.GetUser(ctx)
	if err != nil {
		return 0, err
	}

	return user.Id, err
}

func (a *frontTokenHandle) Remove(ctx context.Context, token string) (err error) {
	a.user = nil
	return a.TokenHandler.Remove(ctx, token)
}
