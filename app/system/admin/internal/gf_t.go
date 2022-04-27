package main

import (
	"fmt"
	"gf-admin/utility"
	"sync"
	"time"
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

func main() {

	fmt.Println(utility.EncryptPassword("admina", "123456"))
	//tt.Set()
	//
	//fmt.Println(tt.m)
	//
	//c := []int{1, 2, 3, 4}
	//
	//fmt.Println(c[:1])
	//d := append(c[:0], c[1:]...)
	//g.Dump(d)
	//g.Dump(d[1])
	//
	//m := map[string]int{
	//	"a": 1,
	//	"b": 2,
	//	"c": 3,
	//}
	//
	//for key, value := range m {
	//	delete(m, "a")
	//	fmt.Println(key, value)
	//}

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		time.Sleep(3 * time.Second)
		wg.Add(-2)
	}()

	go func() {
		wg.Wait()

		fmt.Println(111)
	}()

	wg.Wait()
	fmt.Println(2222)

}
