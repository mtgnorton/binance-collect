package main

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"os"
)

const BYTE = 1 << (10 * iota)

type T struct {
	m map[string][]int
}

var tt = T{
	m: map[string][]int{
		"a": []int{6},
	},
}

func (t *T) Set() {
	t.m["a"] = []int{1, 2, 3}
}

type CustomCode struct {
	code    int
	message string
	detail  interface{}
}

type CustomDetail struct {
	CustomType string
}

// Code returns the integer number of current error code.
func (c *CustomCode) Code() int {
	return c.code
}

// Message returns the brief message for current error code.
func (c *CustomCode) Message() string {
	return c.message
}

// Detail returns the detailed information of current error code,
// which is mainly designed as an extension field for error code.
func (c *CustomCode) Detail() interface{} {
	return c.detail
}

func NewCustomCode(code int, message string, CustomType string) *CustomCode {
	return &CustomCode{
		code:    code,
		message: message,
		detail:  CustomDetail{CustomType: CustomType},
	}
}

func test4() {

}
func test3() {
	test4()
}
func test2() {
	test3()
}
func test1() {
	test2()
}
func ExampleStack() {
	var err error
	err = errors.New("sql error")
	err = gerror.Wrap(err, "adding failed")
	err = gerror.Wrap(err, "api calling failed")
	fmt.Println(gerror.Stack(err))

	// Output:
	// 1. api calling failed
	//     1).  main.main
	//         /Users/john/Workspace/Go/GOPATH/src/github.com/gogf/gf/.example/other/test.go:14
	// 2. adding failed
	//     1).  main.main
	//         /Users/john/Workspace/Go/GOPATH/src/github.com/gogf/gf/.example/other/test.go:13
	// 3. sql error
}
func main() {

	ExampleStack()
	os.Exit(1)

	customCode := NewCustomCode(1, "custom", "custom")
	errorCustomCode := gerror.NewCode(customCode)
	code := gerror.Code(errorCustomCode)
	g.DumpWithType(code)
	fmt.Println(code)
}
