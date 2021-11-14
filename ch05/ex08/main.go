package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file := os.Args[1]
	id := os.Args[2]

	f, _ := os.Open(file)
	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	n := ElementByID(doc, id)
	if n != nil {
		fmt.Printf("id='%s' found in %v: attributes -> %v\n", id, n.Data, n.Attr)
	} else {
		fmt.Printf("id='%s' not found\n", id)
	}
}

func findID(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}
	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode && findID(n, id) {
			node = n
			return false // 走査中止
		}
		return true // 走査続行
	}, nil)
	return node
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}
	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}
