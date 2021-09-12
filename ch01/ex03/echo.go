package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo_without_strings(args []string) error {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
	return nil
}

func echo_with_strings(args []string) error {
	fmt.Println(strings.Join(args[1:], " "))
	return nil
}

func main() {
	args := os.Args
	// 非効率な可能性のあるバージョン
	start := time.Now()
	echo_without_strings(args)
	secs := time.Since(start).Seconds()
	fmt.Printf("非効率な可能性のあるバージョン: %.8fs\n", secs)
	// strings.Joinを使ったバージョン
	start = time.Now()
	echo_with_strings(args)
	secs = time.Since(start).Seconds()
	fmt.Printf("strings.Joinを使ったバージョン: %.8fs\n", secs)
}
