package main

import (
	"ex16/strings"
	"fmt"
)

func main() {
	s := strings.Join(" ", "a", "b", "c")
	fmt.Println(s) // a b c
	s = strings.Join("/", "a", "b", "c")
	fmt.Println(s) // a/b/c
}
