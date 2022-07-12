package main

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gstructs"
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

func main() {

	m := cMain{}
	reflectValue := reflect.ValueOf(m)

	//如果是struct，则无法获取到指针方法,需要构建一个该struct对应的指针Value
	if reflectValue.Kind() == reflect.Struct {
		//reflect.New 返回一个reflect.Value,该Value是一个指向某个类型的新的零值的指针
		newValue := reflect.New(reflectValue.Type())
		newValue.Elem().Set(reflectValue)
		reflectValue = newValue
	}

	//reflectValue.Method返回reflect.Value
	method := reflectValue.Method(0)
	methodType := method.Type()

	//*reflect.rtype(19) "main.cMainHttpInput"
	inArgType := methodType.In(1)

	//reflect.Kind(6) "struct"
	inArgKind := inArgType.Kind()

	var inputObject reflect.Value

	if inArgKind == reflect.Ptr {
		inputObject = reflect.New(inArgType.Elem()).Elem()
	} else {

		//Elem()返回接口或指针指向的值,类型为reflect.Value，如果调用者不是接口或指针，则会panic
		//reflect.Value(2) {
		//    Name: string(0) "",
		//    Port: int(0),
		//}
		inputObject = reflect.New(inArgType).Elem()
	}

	//Addr()返回指向底层值的指针，为reflect.Value类型

	g.DumpWithType(111, inputObject.Addr().Interface())
	//Interface()返回调用者的底层值，类型为interface{}
	inputObjectInterface := inputObject.Interface()

	//gstructs.Fields()返回Pointer指向的struct的所有字段
	//[]gstructs.Field(2) [
	//    gstructs.Field(3) {
	//        Value:    reflect.Value(0) {},
	//        Field:    reflect.StructField(7) {
	//            Name:      string(4) "Name",
	//            PkgPath:   string(0) "",
	//            Type:      *reflect.rtype(6) "string",
	//            Tag:       reflect.StructTag(55) "v:\"required\" name:\"NAME\" arg:\"true\" brief:\"server name\"",
	//            Offset:    uintptr(0),
	//            Index:     []int(1) [
	//                int(1),
	//            ],
	//            Anonymous: bool(false),
	//        },
	//        TagValue: string(0) "",
	//    },
	//    gstructs.Field(3) {
	//        Value:    reflect.Value(11) "<int Value>",
	//        Field:    reflect.StructField(7) {
	//            Name:      string(4) "Port",
	//            PkgPath:   string(0) "",
	//            Type:      *reflect.rtype(3) "int",
	//            Tag:       reflect.StructTag(63) "v:\"required\" short:\"p\" name:\"port\"  brief:\"port of http server\"",
	//            Offset:    uintptr(16),
	//            Index:     []int(1) [
	//                int(2),
	//            ],
	//            Anonymous: bool(false),
	//        },
	//        TagValue: string(0) "",
	//    },
	//]
	argFields, err := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         inputObjectInterface,
		RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
	})

	g.Dump(err)
	//map[string]string(4) {
	//    string("arg"):   string(4) "true",
	//    string("brief"): string(11) "server name",
	//    string("v"):     string(8) "required",
	//    string("name"):  string(4) "NAME",
	//}
	argFieldTagMap := argFields[0].TagMap()

	//
	g.DumpWithType(argFieldTagMap)

	cmd, err := gcmd.NewFromObject(cMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
