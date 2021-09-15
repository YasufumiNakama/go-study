package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	// Get "https://www.zhanqi.tv": net/http: TLS handshake timeout みたいなエラーが起こり得る
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.tmall.com",
		"https://www.baidu.com",
		"https://www.qq.com",
		"https://www.sohu.com",
		"https://www.facebook.com",
		"https://www.taobao.com",
		"https://www.360.cn",
		"https://www.jd.com",

		"https://www.amazon.com",
		"https://www.yahoo.com",
		"https://www.wikipedia.org",
		"https://www.weibo.com",
		"https://www.sina.com.cn",
		"https://www.zoom.us",
		"http://www.xinhuanet.com",
		"https://www.live.com",
		"https://www.netflix.com",
		"https://www.reddit.com",

		"https://www.instagram.com",
		"https://www.microsoft.com",
		"https://www.office.com",
		"https://www.google.com.hk",
		"https://panda.tv",
		"https://www.zhanqi.tv",
		"https://www.alipay.com",
		"https://www.bing.com",
		"https://www.csdn.net",
		"https://www.myshopify.com",

		"https://www.vk.com",
		"https://www.yahoo.co.jp",
		"https://www.bongacams.com",
		"https://login.microsoftonline.com",
		"https://www.naver.com",
		"https://www.twitch.tv",
		"https://www.twitter.com",
		"https://www.okezone.com",
		"https://www.aparat.com",
		"https://www.ebay.com",

		"https://www.amazon.in",
		"https://www.adobe.com",
		"https://www.aliexpress.com",
		"https://www.yy.com",
		"https://www.tianya.cn",
		"https://www.huanqiu.com",
		"https://www.chaturbate.com",
		"https://www.amazon.co.jp",
		"https://www.canva.com",
		"https://www.linkedin.com",
	}
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
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
