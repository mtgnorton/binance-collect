package controller

import (
	"context"
	"gf-admin/app/system/admin/internal/define"
	"gf-admin/app/system/admin/internal/service"

	gfg "github.com/gogf/gf/v2/frame/g"
)

var Global = global{}

type global struct {
}

func (g *global) Logout(ctx context.Context, req *define.AdministratorLogoutReq) (res *define.AdministratorLogoutRes, err error) {
	res = &define.AdministratorLogoutRes{}
	token, err := service.AdminTokenInstance.GetTokenFromRequest(ctx, gfg.RequestFromCtx(ctx))
	if err != nil {
		return
	}
	err = service.AdminTokenInstance.Remove(ctx, token)
	return
}

func (g *global) GetLoggedInfo(ctx context.Context, req *define.AdministratorGetLoggedInfoReq) (res *define.AdministratorGetLoggedInfoRes, err error) {
	res = &define.AdministratorGetLoggedInfoRes{}
	res.AdministratorSummary, err = service.AdminTokenInstance.GetAdministrator(ctx)
	return
}

func (g *global) Routes(ctx context.Context, req *define.RoutesReq) (res *define.RoutesRes, err error) {
	adminId, err := service.AdminTokenInstance.
		GetAdministratorId(ctx)
	if err != nil {
		return
	}
	res = &define.RoutesRes{}
	res.FrontRoutes, err = service.Menu.FrontRoutes(ctx, adminId)
	return
}

//func (g *global) Ws(ctx context.Context, req *define.WsReq) (res define.WsRes, globalErr error) {
//	r := gfg.RequestFromCtx(ctx)
//	ws, err := r.WebSocket()
//	if err != nil {
//		gfg.Log().Error(ctx, err)
//		r.Exit()
//	}
//	for {
//		msgType, msg, err := ws.ReadMessage()
//		if err != nil {
//			err = globalErr
//			return
//		}
//		if err = ws.WriteMessage(msgType, msg); err != nil {
//			err = globalErr
//			return
//		}
//	}
//}
