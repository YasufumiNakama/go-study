package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2}
	b := rotate(a, 2)
	fmt.Printf("rotate(%v, 2) = %v\n", a, b)
}

func rotate(slice []int, k int) []int {
	n := len(slice)
	k %= n
	if k == 0 {
		return slice
	}
	return append(slice[n-k:], slice[:n-k]...)
}
