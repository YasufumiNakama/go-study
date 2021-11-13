package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	count := map[string]int{}
	countElementNode(count, doc)
	for ElementNode, cnt := range count {
		fmt.Printf("%v: %v\n", ElementNode, cnt)
	}
}

func countElementNode(count map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	countElementNode(count, n.FirstChild)
	countElementNode(count, n.NextSibling)
}
