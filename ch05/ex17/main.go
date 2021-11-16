package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file := os.Args[1]

	f, _ := os.Open(file)
	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	images := ElementsByTagName(doc, "img")
	for _, image := range images {
		fmt.Println(image.Data)
	}

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, heading := range headings {
		fmt.Println(heading.Data)
	}

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if n == doc.Data {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		for _, node := range ElementsByTagName(c, name...) {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
