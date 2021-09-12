package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout // for testing

/*
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
*/

func echo(args []string) error {
	for i, arg := range args {
		// fmt.Println(i, arg)
		fmt.Fprintln(out, i, arg) // for testing
	}
	return nil
}

func main() {
	echo(os.Args[1:])
}
