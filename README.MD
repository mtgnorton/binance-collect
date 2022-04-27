# gf-admin

# description:

1. 配置使用的config文件,可选值有(local,dev,prod),默认为prod,可以通过命令行参数(gf.admin.env.file)和环境变量(GF_ADMIN_ENV_FILE)设置，当同时设置时，优先使用命令行参数
    - 使用命令行参数，`go run main.go --gf.admin.env.file=local` 
    - 使用环境变量，`export GF_ADMIN_ENV_FILE=local; go run main.go`,gf命令类似`export GF_ADMIN_ENV_FILE=local; gf run main.go`

# todo

1. swagger 分组问题
2. model 数据模型问题

# 注意事项

1.

当注册路由后，出现以下类似错误`2022-04-13 22:22:57.561 [FATA] duplicated route registry "/@default" at /Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin/app/system/admin/admin.go:81 , already registered at /Users/mtgnorton/Coding/go/src/github.com/mtgnorton/gf-admin/app/system/admin/admin.go:81
`,关键词`/@default`,原因为注册的控制器中包含index方法，index方法会和gf框架默认生成的方法重复，导致报错



