package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	// https://xn--go-hh0g6u.com/pkg/unicode/
	counts := map[string]int{
		"Letter": 0, // unicode.IsLetter
		"Number": 0, // unicode.IsNumber
		"Punct":  0, // unicode.IsPunct
		"Space":  0, // unicode.IsSpace
	}
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["Letter"]++
		}
		if unicode.IsNumber(r) {
			counts["Number"]++
		}
		if unicode.IsPunct(r) {
			counts["Punct"]++
		}
		if unicode.IsSpace(r) {
			counts["Space"]++
		}
	}
	fmt.Printf("分類\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
