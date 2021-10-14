package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("こんにちは")
	reverse(b)
	fmt.Println(string(b)) // はちにんこ
}

func reverse(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	_, size := utf8.DecodeRune(b) // https://xn--go-hh0g6u.com/pkg/unicode/utf8/#DecodeRune
	// sizeだけ左へbを回転させる ex. {0, 1, 2, 3, 4, 5} -> size=2 -> {2, 3, 4, 5, 0, 1}
	rev(b[:size])
	rev(b[size:])
	rev(b)
	return reverse(b[:len(b)-size])
}

func rev(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
