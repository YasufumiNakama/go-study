package main

import (
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	f, _ := os.Open("./sample.html")
	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	images := ElementsByTagName(doc, "img")
	if len(images) != 1 {
		t.Errorf("len(%v) is not 1", images)
	}

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	if len(headings) != 2 {
		t.Errorf("len(%v) is not 2", headings)
	}
}
