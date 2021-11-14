package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "abc $foo $bar xxx"
	fmt.Printf("%s -> %s\n", s, expand(s, strings.ToUpper))
}

func expand(s string, f func(string) string) string {
	/*
		https://ashitani.jp/golangtips/tips_regexp.html
		https://murashun.jp/article/programming/regular-expression.html
	*/
	reg := regexp.MustCompile(`\$\S*`)
	out := reg.ReplaceAllStringFunc(s, func(s string) string { return f(s[1:]) })
	return out
}
