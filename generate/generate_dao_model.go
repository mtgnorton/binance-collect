package main

import (
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func main() {

	table := "notify" //ga_admin_log
	ctx := gctx.New()

	command := "gf gen dao"
	if table != "" {
		//使用local配置文件
		command = `gf gen dao -gf.gcfg.file config-local.toml -t ` + table
	}

	cmd := exec.Command("/bin/zsh", "-c", command)
	var output []byte
	var err error
	if output, err = cmd.Output(); err != nil {
		g.Log().Fatalf(ctx, "执行gf gen dao 错误: %s", err)
	}
	g.Dump(output)

	rootPath := gfile.Pwd()

	Dirs := map[string]string{
		"/app/service/internal/dao":          "/app/dao/",
		"/app/service/internal/dao/internal": "/app/dao/internal/",
		"/app/service/internal/do":           "/app/dto/",
	}
	for tempSource, tempDst := range Dirs {
		_, err = gfile.ScanDirFileFunc(rootPath+tempSource, "", false, func(path string) string {
			dst := rootPath + tempDst + gfile.Basename(path)
			content := gfile.GetContents(path)

			content = gstr.Replace(content, "gf-admin/app/service/internal/dao/internal", "gf-admin/app/dao/internal")

			content = gstr.Replace(content, "package do", "package dto")
			err = gfile.PutContents(dst, content)
			if err != nil {
				g.Log().Fatalf(ctx, "写入文件错误：%s", err)
			}
			g.Log().Infof(ctx, "%s 移动到 %s \n", path, dst)
			return path
		})
		fmt.Println(err)
	}

	err = gfile.Remove(rootPath + "/app/service")
	fmt.Println(err)
}
