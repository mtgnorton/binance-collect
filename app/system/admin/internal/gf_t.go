package main

import "fmt"

type A struct {
	a []int
}

func main() {

	as := A{}

	for _, item := range as.a {
		fmt.Println(item)
	}

}
