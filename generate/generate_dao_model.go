package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"os/exec"
)

func main() {

	table := "" //ga_admin_log
	ctx := gctx.New()

	command := "gf gen dao"
	if table != "" {
		command = `gf gen dao -t ` + table
	}

	cmd := exec.Command("/bin/zsh", "-c", command)
	var output []byte
	var err error
	if output, err = cmd.Output(); err != nil {
		g.Log().Fatalf(ctx, "执行gf gen dao 错误: %s", err)
		return
	}
	g.Dump(output)

	rootPath := gfile.Pwd()

	Dirs := map[string]string{
		"/app/service/internal/dao":          "/app/dao/",
		"/app/service/internal/dao/internal": "/app/dao/internal/",
		"/app/service/internal/dto":          "/app/dto/",
	}

	for tempSource, tempDst := range Dirs {
		gfile.ScanDirFileFunc(rootPath+tempSource, "", false, func(path string) string {
			dst := rootPath + tempDst + gfile.Basename(path)
			content := gfile.GetContents(path)

			content = gstr.Replace(content, "gf-admin/app/service/internal/dao/internal", "gf-admin/app/dao/internal")
			err = gfile.PutContents(dst, content)
			if err != nil {
				g.Log().Fatalf(ctx, "写入文件错误：%s", err)
			}
			g.Log().Infof(ctx, "%s 移动到 %s \n", path, dst)
			return path
		})
	}

	gfile.Remove(rootPath + "/app/service")
}
