package main

import (
	"fmt"
	//"sort"
)

var prereqs = map[string]map[string]bool{ // var prereqs = map[string][]string{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures": true,
		"formal languages": true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string { // func topoSort(m map[string][]string) []string {
	/* スライスの代わりにマップを使う */
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool) // var visitAll func(items []string)

	visitAll = func(items map[string]bool) { // visitAll = func(items []string) {
		for item := range items { // for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool) // var keys []string
	for key := range m {
		keys[key] = true // keys = append(keys, key)
	}

	// sort.Strings(keys)
	visitAll(keys)
	return order
}