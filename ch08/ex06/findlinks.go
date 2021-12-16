// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"
	// "os"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var maxDepth = flag.Int("depth", 3, "max depth of links")

type depthList struct {
	depth int
	list  []string
}

func crawl(depth int, url string) *depthList {
	if depth > *maxDepth {
		return &depthList{depth + 1, nil}
	}
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return &depthList{depth + 1, list}
}

//!-sema

//!+
func main() {
	flag.Parse()
	// worklist := make(chan []string)
	worklist := make(chan *depthList)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	// go func() { worklist <- os.Args[1:] }()
	go func() { worklist <- &depthList{0, flag.Args()} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		dList := <-worklist
		for _, link := range dList.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(depth int, link string) {
					worklist <- crawl(depth, link)
				}(dList.depth, link)
			}
		}
	}
}

//!-
