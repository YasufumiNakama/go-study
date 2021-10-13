package main

import (
	"fmt"
	"strings"
)

func main() {
	sa := "あいうえお"
	sb := "うおあいえ"
	fmt.Printf("sa: %s\n", sa)
	fmt.Printf("sb: %s\n", sb)
	if anagram(sa, sb) {
		fmt.Println("it's anagram")
	} else {
		fmt.Println("it's NOT anagram")
	}
}

func anagram(sa, sb string) bool {
	for _, ra := range sa {
		b := strings.ContainsRune(sb, ra)
		if !b {
			return false
		}
		sb = strings.Replace(sb, string(ra), "", 1)
	}
	return sb == ""
}
