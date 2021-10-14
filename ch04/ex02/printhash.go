package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha = flag.String("sha", "sha256", "select sha from sha256, sha384, sha512")

func main() {
	flag.Parse()
	if *sha == "sha256" || *sha == "sha384" || *sha == "sha512" {
		fmt.Printf("標準入力の%sハッシュを表示します\n", *sha)
	} else {
		fmt.Fprintf(os.Stderr, "%s: invalid sha, select sha from sha256, sha384, sha512\n", *sha)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d := scanner.Bytes()
		if *sha == "sha256" {
			fmt.Printf("sha256: %x\n", sha256.Sum256(d))
		} else if *sha == "sha384" {
			fmt.Printf("sha384: %x\n", sha512.Sum384(d))
		} else if *sha == "sha512" {
			fmt.Printf("sha512: %x\n", sha512.Sum512(d))
		}
	}
}
