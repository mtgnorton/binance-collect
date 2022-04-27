package shared

import (
	"context"
	"fmt"
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"strings"
	"time"
)

const (
	CacheModeCache = "cache"
	CacheModeRedis = "redis"
)

type TokenFrame struct {
	UserKey     string
	UUID        string
	Token       string
	CreateTime  int64
	RefreshTime int64
	Data        interface{}
}

func (t *TokenFrame) GetUerKey() string {
	return t.UserKey
}
func (t *TokenFrame) GetUUID() string {
	return t.UUID
}

func (t *TokenFrame) GetToken() string {
	return t.Token
}
func (t *TokenFrame) GetData() interface{} {
	return t.Data
}

func (t *TokenFrame) String() string {
	return fmt.Sprintf("\n"+
		"UserKey:%s\n"+
		"UUID:%s\n"+
		"Token:%s\n"+
		"CreateTime:%s\n"+
		"RefreshTime:%s\n"+
		"Data:%s\n", t.UserKey, t.UUID, t.Token, gconv.String(t.CreateTime), gconv.String(t.RefreshTime), t.Data)

}

type TokenHandler struct {
	CacheMode  string
	CacheKey   string
	Timeout    int
	MaxRefresh int
	MultiLogin bool
	EncryptKey []byte
	InitFinish bool
}

func (t *TokenHandler) Middleware(group *ghttp.RouterGroup) {
	group.Middleware(func(r *ghttp.Request) {
		g.Log("token").Debug(r.Context(), "token中间件开始执行，token处理器开始获取header头")

		err := t.InitUser(r)
		if err != nil {
			response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		}
		r.Middleware.Next()
	})
}

func (t *TokenHandler) InitUser(r *ghttp.Request) error {
	ctx := r.Context()
	if !t.InitFinish {
		t.Init(ctx)
	}
	token, err := t.GetTokenFromRequest(ctx, r)
	if err != nil {
		return err
	}
	tf, err := t.Validate(ctx, token)
	if err != nil {
		return err
	}

	Context.SetUser(ctx, tf.Data) //将用户数据注入到上下文中

	return nil
}

func (t *TokenHandler) Init(ctx context.Context) {

	if t.CacheMode == "" {
		t.CacheMode = CacheModeCache
	}
	if t.CacheKey == "" {
		t.CacheKey = "administrator_token_"
	}
	if string(t.EncryptKey) == "" {
		t.EncryptKey = []byte("noworldcanexpressmywholehearted1")
	}
	if t.Timeout == 0 {
		t.Timeout = 10 * 24 * 60 * 60 * 1000
	}
	if t.MaxRefresh == 0 {
		t.MaxRefresh = t.Timeout / 2
	}
	t.MaxRefresh = gconv.Int(gtime.Now().TimestampMilli()) + t.MaxRefresh

	t.InitFinish = true
	g.Log("token").Debug(ctx, "token init finish")
}

/*根据userKey和uuid生成token,并将相关数据持久化*/
func (t *TokenHandler) GenerateAndSaveData(ctx context.Context, userKey string, data interface{}) (token string, err error) {
	tf, err := t.encrypt(ctx, userKey)
	if err != nil {
		return "", err
	}
	if t.MultiLogin { //支持多点登录，并且已经登录直接将老的token返回
		r, err := t.GetData(ctx, userKey)

		if r.GetToken() != "" && err == nil {
			g.Log("token").Debugf(ctx, "支持多点登录，直接返回：%s", r.GetToken())
			return r.GetToken(), nil
		}
	}
	cacheKey := t.CacheKey + userKey

	tf.Data = data
	tf.CreateTime = gtime.Now().TimestampMilli()
	tf.RefreshTime = gtime.Now().TimestampMilli() + gconv.Int64(t.MaxRefresh)
	err = t.cacheSet(ctx, cacheKey, tf)
	return tf.GetToken(), err
}

/*根据userkey更新缓存里面的用户相关数据,token的创建时间和过期时间等不会改变*/
func (t *TokenHandler) UpdateData(ctx context.Context, userKey string, data interface{}) (err error) {
	cacheKey := t.CacheKey + userKey
	tf, err := t.cacheGet(ctx, cacheKey)
	if err != nil {
		return
	}
	tf.Data = data
	err = t.cacheSet(ctx, cacheKey, tf)
	return
}

/*验证token有效性*/
func (t *TokenHandler) Validate(ctx context.Context, token string) (tf TokenFrame, err error) {

	tf = TokenFrame{}
	if string(token) == "" {
		return tf, gerror.NewCode(gcode.CodeValidationFailed, "validate Token empty")
	}
	tf1, err := t.decrypt(ctx, token)

	if err != nil {
		return tf1, err
	}
	tf2, err := t.GetData(ctx, tf1.GetUerKey())
	if err != nil {
		return tf2, err
	}

	if tf1.GetUUID() != tf2.GetUUID() {
		return tf2, gerror.New("user Token auth error")
	}
	return tf2, nil

}

/*从header的Authorization获取token*/
func (t *TokenHandler) GetTokenFromRequest(ctx context.Context, r *ghttp.Request) (token string, err error) {
	token = r.GetQuery("token").String()

	if token != "" {
		//g.Log("token").Debugf(ctx, "通过url参数获取到的token为:%s", token)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		g.Log("token").Debugf(ctx, "通过header头获取到的token为:%s", authHeader)
		if len(parts) != 2 || gstr.Trim(parts[0]) != "Bearer" || gstr.Trim(parts[1]) == "" {
			return "", gerror.New("get Token from request error")
		}
		token = gstr.Trim(parts[1])
		return token, nil
	}

	return "", gerror.New("please login ")

}

/*删除token相关数据*/
func (t *TokenHandler) Remove(ctx context.Context, token string) (err error) {
	r, err := t.decrypt(ctx, token)
	if err != nil {
		return err
	}
	err = t.cacheRemove(ctx, t.CacheKey+r.UserKey)
	return err
}

/*根据userKey 获取data*/
func (t *TokenHandler) GetData(ctx context.Context, userKey string) (tf TokenFrame, err error) {
	cacheKey := t.CacheKey + userKey
	tf, err = t.cacheGet(ctx, cacheKey)
	if err != nil {
		return tf, err
	}
	nowTime := gtime.Now().TimestampMilli()
	refreshTime := tf.RefreshTime
	if gconv.Int64(refreshTime) == 0 || nowTime > gconv.Int64(refreshTime) {
		tf.CreateTime = nowTime
		tf.RefreshTime = nowTime + gconv.Int64(t.MaxRefresh)
		g.Log("Token").Debug(ctx, "Token refresh", tf)
		err = t.cacheSet(ctx, cacheKey, tf)
		if err != nil {
			return tf, err
		}
	}
	return tf, nil
}

/*根据userKey和uuid生成token*/
func (t *TokenHandler) encrypt(ctx context.Context, userKey string, uuid ...string) (tf TokenFrame, err error) {
	tf = TokenFrame{UserKey: userKey}
	if userKey == "" {
		return tf, gerror.New("encrypt UserKey empty")
	}
	if len(uuid) == 0 || uuid[0] == "" {
		tf.UUID, err = gmd5.Encrypt(grand.Letters(10))
		if err != nil {
			g.Log("Token").Errorf(ctx, "UUID generate error,error following %s", err)
			return tf, err
		}
	}
	tokenTemp := userKey + "_" + tf.GetUUID()
	token, err := gaes.Encrypt([]byte(tokenTemp), t.EncryptKey)
	if err != nil {
		return tf, gerror.Wrap(err, "gaes encrypt Token error")
	}
	tf.Token = string(gurl.Encode(string(token)))

	return tf, nil
}

/*将token解密为userKey和uuid*/
func (t *TokenHandler) decrypt(ctx context.Context, token string) (tf TokenFrame, err error) {
	tf = TokenFrame{
		Token: token,
	}
	if tf.GetToken() == "" {
		return tf, gerror.New("decrypt Token empty")
	}
	tokenBase64, err := gurl.Decode(token)
	if err != nil {
		return tf, err
	}
	tokenDecrypted, err := gaes.Decrypt([]byte(tokenBase64), t.EncryptKey)
	if err != nil {
		return tf, err
	}
	tokenComponents := gstr.Split(string(tokenDecrypted), "_")
	if len(tokenComponents) < 2 {
		return tf, gerror.New("decrypt token length error")
	}
	tf.UserKey = tokenComponents[0]
	tf.UUID = tokenComponents[1]
	tf.Token = token
	return
}

func (t *TokenHandler) cacheGet(ctx context.Context, key string) (tf TokenFrame, err error) {
	tf = TokenFrame{}
	var valueVar *g.Var
	switch t.CacheMode {
	case CacheModeRedis:
		valueVar, err = g.Redis().Do(ctx, "get", key)
		if err != nil {
			return tf, gerror.Wrap(err, "token get redis cache error")
		}
		if valueVar.IsNil() || valueVar.IsEmpty() {
			return tf, gerror.New("用户未登录")
		}

	default:
		valueVar, err = gcache.Get(ctx, key)
		if err != nil {
			return tf, gerror.Wrap(err, "Token get gcache cache error")
		}

	}
	err = valueVar.Scan(&tf)
	if err != nil {
		return tf, gerror.Wrap(err, "Token redis cache scan error")
	}
	return tf, err

}
func (t *TokenHandler) cacheSet(ctx context.Context, key string, value TokenFrame) (err error) {
	switch t.CacheMode {
	case CacheModeRedis:

		if t.Timeout == 0 {
			_, err = g.Redis().Do(ctx, "set", key, value)
		} else {
			_, err = g.Redis().Do(ctx, "setex", key, t.Timeout/1000, value)
		}
		if err != nil {
			return gerror.Wrap(err, "redis set error")
		}
	default:
		err = gcache.Set(ctx, key, value, gconv.Duration(t.Timeout)*time.Microsecond)
	}
	return err
}

func (t *TokenHandler) cacheRemove(ctx context.Context, key string) (err error) {

	switch t.CacheMode {
	case CacheModeRedis:
		_, err = g.Redis().Do(ctx, "del", key)
		if err != nil {
			return gerror.Wrap(err, "redis del error")
		}
	default:
		_, err = gcache.Remove(ctx, err)
	}
	return err
}
