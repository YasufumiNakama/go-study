package main

import (
	"fmt"
	"os"
	"strings"
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
	// 符号付きの場合
	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + comma(s[1:])
	}
	// 浮動小数点数の場合
	if i := strings.IndexByte(s, '.'); i > 0 {
		return comma(s[:i]) + s[i:]
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
