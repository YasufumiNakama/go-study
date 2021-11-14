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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

/* see view-source:https://golang.org/ as sample */
var DataKeyMap = map[string]string{
	"a": "href",
	"link": "href", // スタイルシート
	"script": "src", // スクリプト
	"img": "src", // 画像
	"iframe": "src", // 動画
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		for Data, Key := range DataKeyMap {
			if n.Data == Data {
				for _, a := range n.Attr {
					if a.Key == Key {
						links = append(links, a.Val)
					}
				}
			}
		}
	}
	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}