# gf-admin

# description:

1. 配置文件： 可选值有(local,dev,prod),默认为prod,可以通过命令行参数(gf.admin.env.file)和环境变量(GF_ADMIN_ENV_FILE)设置，当同时设置时，优先使用命令行参数
    - 使用命令行参数，`go run main.go --gf.admin.env.file=local`
    - 使用环境变量，`export GF_ADMIN_ENV_FILE=local; go run main.go`,gf命令类似`export GF_ADMIN_ENV_FILE=local; gf run main.go`
2. 路由命名规范：如下
    ```
   增删该查相关，以角色举例
    角色列表： RoleListReq,RoleListRes-> role-list
    角色添加: RoleStoreReq,RleStoreRes-> role-store
    角色详情：RoleInfoReq,RoleInfoRes ->role-info
    角色更新：RoleUpdateReq,RoleUpdateRes -> role-update
    角色删除：RoleDestroyReq,RoleDestroyRes-> role-destroy

   get/post相关，以登录举例
   登录get: LoginInfoReq,LoginInfoRes -> login-info
   登录post: LoginReq,LoginRes -> login
    ```
3. 错误处理
   错误会在调用栈中依次返回，最原始的错误需要被gerror.New封装后再返回，api层需要使用gerror.Wrap封装返回给用户的错误消息,错误的堆栈会被记录到默认日志中,使用举例如下：
    ```
    func c3() error {
	err1 := errors.New("sql error")
    return gerror.New(err1.Error()) //使用gerror.New封装,为了打印出最原始的错误堆栈
    }

    func c2() error {
    return c3()
   }

   func c1() error(){
    err := c2()
   if err!=nil{
    gerror.Wrap(err,"数据库错误")//错误信息展示给用户
   }
   }
    ```

# todo

1. 前台用户权限
2. 用户活跃度计算
3. 广告管理
4. 帮助文章管理
5. 主题权重计算



# 前台接口
1. 主题
2. 创建主题
   - hook
   - 扣除货币

3. 编辑主题
4. 移动主题
5. 主题详情
6. 附言
7. 主题列表
   - 用户创建主题列表
   - 首页主题列表(排序规则)
   - 节点主题列表
8. 感谢主题
9. 收藏主题
10. 屏蔽主题

# 注意事项

1.

当注册路由后，出现以下类似错误`2022-04-13 22:22:57.561 [FATA] duplicated route registry "/@default" at /Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin/app/system/admin/admin.go:81 , already registered at /Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin/app/system/admin/admin.go:81
`,关键词`/@default`,原因为注册的控制器中包含index方法，index方法会和gf框架默认生成的方法重复，导致报错



