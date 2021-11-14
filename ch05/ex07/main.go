package main

import (
	"fmt"
	"io"
	//"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var writer io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	/*
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
	*/
	f, _ := os.Open(url)
	doc, err := html.Parse(f)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// 個々の要素の属性
		var attributes string
		for _, a := range n.Attr {
			attributes += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		// 要素が子を持たない場合: <img></img> -> <img/>
		if n.FirstChild == nil {
			fmt.Fprintf(writer, "%*s<%s%s/>\n", depth*2, "", n.Data, attributes)
		} else {
			fmt.Fprintf(writer, "%*s<%s%s>\n", depth*2, "", n.Data, attributes)
			depth++
		}
	} else if n.Type == html.TextNode {
		// テキストノード
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			fmt.Fprintf(writer, "%*s%s\n", depth*2, "", s)
		}
	} else if n.Type == html.CommentNode {
		// コメントノード
		fmt.Fprintf(writer, "%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	// 要素が子を持たない場合: <img></img> -> <img/>
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Fprintf(writer, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
