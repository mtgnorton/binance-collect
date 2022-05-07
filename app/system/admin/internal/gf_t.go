package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

type cMain struct {
	g.Meta `name:"main" brief:"start http server"`
}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
	Name   string `v:"required" name:"NAME" arg:"true" brief:"server name"`
	Port   int    `v:"required" short:"p" name:"port"  brief:"port of http server"`
}
type cMainHttpOutput struct{}

func (c *cMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainHttpOutput, err error) {
	s := g.Server(in.Name)
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello world")
	})
	s.SetPort(in.Port)
	s.Run()
	return
}

type Runnable = interface {
	Run()
}

func main() {

	cm := cMain{}
	rt := reflect.TypeOf(cm)
	rv := reflect.ValueOf(cm)
	g.DumpWithType(rt.Kind())
	g.DumpWithType(rv.Kind())

	cmp := &cMain{}
	rtp := reflect.TypeOf(cmp)
	rvp := reflect.ValueOf(cmp)
	g.DumpWithType(rtp.Kind(), rtp.Elem())
	g.DumpWithType(rvp.Kind(), rvp.Elem(), rvp.Elem().Kind())

	fmt.Println("======")
	s := make([]string, 0)
	rtp = reflect.TypeOf(s)
	rvp = reflect.ValueOf(s)
	g.DumpWithType(rtp.Kind(), rtp.Elem())

	fmt.Println("====")
	var i io.Writer
	i = os.Stdout
	rtp = reflect.TypeOf(i)
	rvp = reflect.ValueOf(i)
	g.DumpWithType(rtp.Kind(), rtp.Elem())
	g.DumpWithType(rvp.Kind(), rvp.Elem(), rvp.Elem().Kind())

	//g.DumpWithType(rvp.Kind(), rvp.Elem(), rvp.Elem().Kind())
	os.Exit(1)

	cmd, err := gcmd.NewFromObject(cMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
