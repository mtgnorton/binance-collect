package shared

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"
)

var WsService = wsService{
	wsUsersMap: make(map[uint][]*wsUser),
	timeout:    10,
	userCount:  0,
	connCount:  0,
}

type wsService struct {
	wsUsersMap map[uint][]*wsUser //使用slice是因为可能需要多设备登录
	timeout    int64              //超时时间，连接多久没有进行交互后移除
	sync.RWMutex
	userCount int64 // 在线用户数量统计，包含一个账号多处登录
	connCount int64 //连接数量统计
}

//判断连接是否仍然存活
func init() {
	gtimer.AddSingleton(gctx.New(), time.Second*5, func(ctx context.Context) {

		fmt.Println("clear开始运行")
		for _, users := range WsService.wsUsersMap {
			for index, user := range users {
				if gtime.Timestamp()-user.lastSendTime > WsService.timeout || user.close {
					WsService.RemoveUser(user.Id, index)
					g.Log("ws").Debugf(ctx, "移除用户id为：%d，index为：%d的用户", user.Id, index)
				}
			}
		}

		g.Log("ws").Debugf(ctx, "当前已经连接数量为:%d,用户数量为:%d", WsService.ConnCount(), WsService.UserCount())
	})
}

// WsUser WebSocket 用户标示
type wsUser struct {
	Id                uint
	Conn              *ghttp.WebSocket
	writeMessagesChan chan interface{} //ws消息队列，当多个进程同时向同一个conn发送消息，并且消息量大时，可能会有并发问题，使用chan改为串行
	lastSendTime      int64            //最后一次推送消息的时间,超过10s移除用户
	close             bool             //清理失活用户连接有两种可能性，第一种通过定时任务主动检测到用户断开，可以进行主动清理，第二种没有检测到，此时用户的conn已经不可用，但是后端感知不到，在WriteLoop函数抛错后才能感知到,此时将用户标记为close，等待定时任务清理
	userAgent         string           //用户浏览器标识
}

// 广播信息
func (ws *wsService) Broadcast(message *WsMessage) {
	for _, users := range ws.wsUsersMap {
		for _, user := range users {
			user.Write(message)
		}
	}
}

// 获取用户的数量
func (ws *wsService) UserCount() int64 {
	return atomic.LoadInt64(&ws.userCount)
}

//统计连接的数量
func (ws *wsService) ConnCount() int64 {
	return atomic.LoadInt64(&ws.connCount)

	//var count int
	//
	//for _, users := range ws.wsUsersMap {
	//
	//	count += len(users)
	//}
	//return count
}

//获取所有用户
func (ws *wsService) GetUsers() map[uint][]*wsUser {
	return ws.wsUsersMap
}

//判断某个用户是否存在
func (ws *wsService) ExistUser(userId uint) bool {
	return len(ws.wsUsersMap[userId]) > 0
}

//获取某个用户
func (ws *wsService) GetUser(userId uint) []*wsUser {
	return ws.wsUsersMap[userId]
}

//判断userId&&agent 是否是唯一的
func (ws *wsService) isUserAndAgentUnique(userId uint, userAgent string) bool {

	isUnique := true
	for _, existUser := range ws.wsUsersMap[userId] {

		if userAgent == existUser.userAgent {
			isUnique = false
		}
	}

	return isUnique
}

//添加用户,返回用户在slice中的索引位置
func (ws *wsService) AddUser(user *wsUser) (index int) {

	ws.Lock()
	defer ws.Unlock()
	userId := user.Id

	if len(ws.wsUsersMap[userId]) == 0 {
		ws.wsUsersMap[userId] = make([]*wsUser, 0)
	}
	if ws.isUserAndAgentUnique(userId, user.userAgent) {
		ws.userCount++
		//atomic.AddInt64(&ws.userCount, 1)
	}
	//atomic.AddInt64(&ws.connCount, 1)

	ws.connCount++
	ws.wsUsersMap[userId] = append(ws.wsUsersMap[userId], user)

	go user.WriteLoop()
	return len(ws.wsUsersMap[userId]) - 1
}

//移除用户
//index为该用户在slice里面的索引位置
func (ws *wsService) RemoveUser(userId uint, indexOption ...int) bool {
	ws.Lock()
	defer ws.Unlock()

	users := ws.GetUser(userId)

	if len(users) == 0 {
		return true
	}
	var index int

	if len(indexOption) > 0 { //移除单个用户
		index = indexOption[0]
		if ws.isUserAndAgentUnique(userId, users[index].userAgent) {
			//atomic.AddInt64(&ws.userCount, -1)
			ws.userCount--
		}
		//atomic.AddInt64(&ws.connCount, -1)
		ws.connCount--
		//users[index].Clear()
		users = append(users[:index], users[index+1:]...)
		ws.wsUsersMap[userId] = users

		fmt.Println("ssssssfsdfs")
		for _, user := range users { //移除整个slice
			user.Clear()
		}
		delete(ws.wsUsersMap, userId)
	}

	return true

}

func NewWsUser(conn *ghttp.WebSocket, userAgent string, userIdOption ...uint) *wsUser {
	var userId uint
	if len(userIdOption) > 0 {
		userId = userIdOption[0]
	}
	user := &wsUser{
		Id:                userId,
		Conn:              conn,
		userAgent:         userAgent,
		writeMessagesChan: make(chan interface{}, 500),
	}
	user.UpdateLastSendTime()
	return user
}

func (user *wsUser) SetUserId(id uint) {
	user.Id = id
}

func (user *wsUser) UpdateLastSendTime() {
	user.lastSendTime = gtime.Timestamp()
}

func (user *wsUser) WriteLoop() {
	for msg := range user.writeMessagesChan {
		if user.close {
			return
		}
		err := user.Conn.WriteJSON(msg)
		if err != nil {
			g.Log("ws").Debugf(gctx.New(), "用户id%v连接已关闭", user.Id)
			user.close = true //等待定时任务进行清理
			return
		}
	}
}

func (user *wsUser) Write(data *WsMessage) {
	if user.close {
		return
	}
	user.writeMessagesChan <- data
}

func (user *wsUser) Read() (messageType int, p []byte, err error) {
	return user.Conn.ReadMessage()
}

func (user *wsUser) Clear() {
	defer func() {
		err := recover()

		if err != nil {
			g.Log("ws").Debugf(gctx.New(), "recover error,%s", err)
		}

	}()
	user.close = true

	err := user.Conn.Close()
	if err != nil {
		g.Log("ws").Debugf(gctx.New(), "close ws  conn error,%s", err)

	}

	close(user.writeMessagesChan)
}

const (
	WsMessageTypeHeart   = "heart"
	WsMessageTypeSystem  = "system"
	WsMessageTypeBinance = "binance"
	WsMessageTypeError   = "error"
	WsMessageTypeDebug   = "debug"
)

//接受和发送的消息格式
type WsMessage struct {
	Type    string      `json:"type"`              // 可选值：heart(心跳包), system(系统信息),error(错误),debug(调试)
	Message string      `json:"message,omitempty"` // 提示信息
	Data    interface{} `json:"data,omitempty"`    // 返回数据(业务接口定义具体数据结构)
}

func TransferWsMessage(p []byte) (m *WsMessage, err error) {
	m = &WsMessage{}
	err = gconv.Scan(p, m)

	if err != nil {
		return m, gerror.New("输入消息转换失败")
	}
	return
}
