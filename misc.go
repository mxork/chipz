package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func ck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
