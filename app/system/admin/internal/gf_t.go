package main

import "fmt"

type A struct {
	a []int
}

func main() {

	var a [5]int

	a[0] = 1

	fmt.Println(a)

	b := new([]int)

	(*b)[0] = 1

	fmt.Println(b)

}
