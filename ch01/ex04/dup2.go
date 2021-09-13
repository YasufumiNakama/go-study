// Goで多次元マップ（複数のキーからなるマップ）を実現したいときにはどうするか: https://qiita.com/ruiu/items/476f65e7cec07fd3d4d7
// Iterate through the fields of a struct in Go: https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type fileline struct {
	file, line string
}

func main() {
	counts := make(map[fileline]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for fileline, n := range counts {
		v := reflect.ValueOf(fileline)
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, v.Field(1), v.Field(0))
		}
	}
}

func countLines(f *os.File, counts map[fileline]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[fileline{arg, input.Text()}]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
