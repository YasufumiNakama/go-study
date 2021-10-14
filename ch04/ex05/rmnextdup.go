package main

import (
	"fmt"
)

func main() {
	a := []string{"1", "1", "2", "2"}
	a = rmnextdup(a) // ["1", "2"]
	fmt.Println(a)
}

func rmnextdup(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	previous := strings[0]
	i := 1
	for _, s := range strings {
		if s != previous {
			previous = s
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
