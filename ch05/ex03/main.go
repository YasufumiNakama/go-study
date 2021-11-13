package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	printTextNode(os.Stdout, doc)
}

func printTextNode(w io.Writer, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		// ウェブブラウザでは内容が表示されない<script>要素と<style>要素の中は調べない
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			io.WriteString(w, n.Data)
		}
	}
	printTextNode(w, n.FirstChild)
	printTextNode(w, n.NextSibling)
}
