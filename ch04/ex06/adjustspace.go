package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := []byte("xx   XX   xx")
	a = adjustspace(a) // []byte("xx XX xx")
	fmt.Println(string(a))
}

func adjustspace(slice []byte) []byte {
	i := 1
	for j, _ := range slice {
		if j == 0 {
			continue
		}
		if unicode.IsSpace(rune(slice[j])) {
			if unicode.IsSpace(rune(slice[j-1])) {
				continue
			}
			slice[i] = ' '
			i++
		} else {
			slice[i] = slice[j]
			i++
		}
	}
	return slice[:i]
}
