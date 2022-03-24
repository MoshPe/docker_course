package main

import (
	"fmt"
)

func main() {
	var arr []int
	arr = make([]int, 7)
	for i := range arr {
		arr[i] = i
	}

	for _, item := range arr {
		fmt.Println(item)
	}
}
