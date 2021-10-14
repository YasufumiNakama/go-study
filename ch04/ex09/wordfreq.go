package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters

	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	fmt.Printf("word\tcount\n")
	for s, n := range counts {
		fmt.Printf("%s\t%d\n", s, n)
	}
}
