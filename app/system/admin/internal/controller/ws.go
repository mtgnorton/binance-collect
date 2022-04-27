package controller

import (
	"context"
	"fmt"
	"gf-admin/app/model"
	"gf-admin/app/shared"
	"gf-admin/app/system/admin/internal/service"
	"gf-admin/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"
)

var Ws = ws{}

type ws struct {
}

func (w *ws) Ws(r *ghttp.Request) {

	ws, err := r.WebSocket()
	if err != nil {
		g.Log("ws").Error(r.Context(), err)
		r.Exit()
	}

	userAgent := r.Header.Get("User-Agent")
	wsUser := service.NewWsUser(ws, userAgent)

	var administrator *model.AdministratorSummary

	if administrator, err = w.auth(r); err != nil {
		wsUser.Write(&service.WsMessage{
			Type:    "error",
			Message: err.Error(),
		})
		g.Log("ws").Errorf(r.Context(), "ws授权验证失败：%s", err)
		r.Exit()
	}
	wsUser.SetUserId(administrator.Id)
	service.WsService.AddUser(wsUser)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err, 5555)
			return
		}
		wm, err := service.TransferWsMessage(msg)

		wsUser.UpdateLastSendTime()

		if wm.Type == service.WsMessageTypeHeart {
			wsUser.Write(&service.WsMessage{
				Type: service.WsMessageTypeHeart,
				Data: g.Map{
					"pong": gtime.Timestamp(),
				},
			})
			continue
		}

		fmt.Println("接收到的用户消息为：", wm)
		if err != nil {
			wsUser.Write(&service.WsMessage{
				Type:    "error",
				Message: err.Error(),
			})
			continue
		}
		wsUser.Write(wm)
	}
}

func (w *ws) auth(r *ghttp.Request) (administrator *model.AdministratorSummary, err error) {
	g.Log("auth").Debug(r.Context(), "ws是否登录验证开始执行")

	customCtx := &model.Context{
		Data: make(g.Map),
	}
	shared.Context.Init(r, customCtx)

	administrator = &model.AdministratorSummary{}

	err = service.AdminTokenInstance.LoadConfig().InitUser(r)

	fmt.Println(err, 66666)
	if err != nil {
		return administrator, err
	}

	administrator, err = service.AdminTokenInstance.GetAdministrator(r.Context())
	if err != nil {
		return
	}
	if administrator.Id == 0 {
		return administrator, gerror.New("未登录或会话已过期，请您登录后再继续")
	}

	return
}

func (ws *ws) MonitorSystem(ctx context.Context) {
	gtimer.AddSingleton(ctx, 1*time.Second, func(ctx context.Context) {

		connAmount := service.WsService.ConnCount()
		if connAmount > 0 {
			memoryInfo, err := utility.GetMemoryInfo()
			if err != nil {
				return
			}
			cpuInfo, err := utility.GetCpuInfo()
			if err != nil {
				return
			}

			service.WsService.Broadcast(&service.WsMessage{
				Type: service.WsMessageTypeSystem,
				Data: g.Map{
					"cpu":                 cpuInfo,
					"memory":              memoryInfo,
					"administratorAmount": service.WsService.UserCount(),
				},
			})
		}
	})
}
