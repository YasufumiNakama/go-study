// Go write file: https://zetcode.com/golang/writefile/
// Go で日付をフォーマットする場合は "2006-01-02" と書く: https://kakakakakku.hatenablog.com/entry/2016/03/28/001145

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	filename := time.Now().Format("2006-01-02_15:04:05") + ".txt"
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for range os.Args[1:] {
		msg := <-ch // receive from channel ch
		fmt.Println(msg)
		_, err = f.WriteString(fmt.Sprintf("%s\n", msg))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	msg := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	_, err = f.WriteString(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
