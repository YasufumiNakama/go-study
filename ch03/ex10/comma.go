package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	b := []byte(s)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(b[i])
	}
	return buf.String()
}
