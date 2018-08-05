// --------------------------------------------------------------
// author:   hoojo
// email:    hoojo_@126.com
// github:   https://github.com/hooj0
// create:   2018-07-30
// copyright by hoojo@2018
// --------------------------------------------------------------



/* ---------------------------------------------------------------
mutex web crawler
==================================================================
web 爬虫
------------------------------------------------------------------
爬虫抓取URL，存放在map对象中，利用mutex锁进行互斥，保证map数据不被共享。
------------------------------------------------------------------*/

package main

import (
	"fmt"
	"sync"
	)


type Cache struct {
	caches map[string]bool
}

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, resources chan string, quit chan string) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。

	// 下面并没有实现上面两种情况：

	if depth <= 0 {
		quit <- "q"
		return
	}

	// println("url: ", url, ", cache: ", c.caches[url])
	if c.caches[url] == true {
		// resources <-  "cache ->" + url
		quit <- "q"
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("-->>", err)
		quit <- "q"
		return
	}

	mutex.Lock()
	c.caches[url] = true
	mutex.Unlock()

	resources <- fmt.Sprintf("found: %s %q", url, body)

	end := make(chan string)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, resources, end)
	}


	for i := 0; i < len(urls); i++ {
		<- end
	}

	quit <- "quit"
	return
}

var c = &Cache{ caches: make(map[string]bool) }
var mutex = sync.Mutex{}

func main() {

	resources := make(chan string)
	quit := make(chan string)

	go Crawl("https://golang.org/", 4, fetcher, resources, quit)

	for {

		select {
			case url := <- resources:
				fmt.Println(url)
			case <- quit:
				return

		}
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
