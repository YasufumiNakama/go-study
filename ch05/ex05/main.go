package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("words: %v, images: %v\n", words, images)
	}
}

// CountWordsAndImages は HTML ドキュメントに対する HTTP GET リクエストを url へ行い、そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		in := strings.NewReader(n.Data)
		scanner := bufio.NewScanner(in)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	word, image := countWordsAndImages(n.FirstChild)
	words += word
	images += image
	word, image = countWordsAndImages(n.NextSibling)
	words += word
	images += image
	return
}
