package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	url   string
	sizes int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{url: url, sizes: len(body)}
}

func main() {
	page := make(chan Page)
	urls := []string{"https://zhuanlan.zhihu.com/p/662142510",
		"https://www.iana.org/help/example-domains",
		"https://zhuanlan.zhihu.com"}
	for _, str := range urls {
		go responseSize(str, page)
	}
	for _, _ = range urls {
		fmt.Println(<-page)
	}

}
