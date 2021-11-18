package main

import (
	"bufio"
	"fmt"
)

type WordCounter int
type LineCounter int

func main() {
	var x LineCounter
	x.Write([]byte("hello\nworld\n!"))
	fmt.Println(x) // 3

	var y WordCounter
	y.Write([]byte("hello world\n!"))
	fmt.Println(y) // 3
}

func (c *LineCounter) Write(p []byte) (int, error) {
	/* https://xn--go-hh0g6u.com/pkg/bufio/#ScanLines */
	n := len(p)
	for len(p) > 0 {
		advance, token, err := bufio.ScanLines(p, true)
		if err == nil && token != nil {
			*c++
		}
		p = p[advance:]
	}
	return n, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	/* https://xn--go-hh0g6u.com/pkg/bufio/#ScanWords */
	n := len(p)
	for len(p) > 0 {
		advance, token, err := bufio.ScanWords(p, true)
		if err == nil && token != nil {
			*c++
		}
		p = p[advance:]
	}
	return n, nil
}
