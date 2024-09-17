package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
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
	fmt.Println(len(body))
}

func main() {
	go responseSize("https://zhuanlan.zhihu.com/p/662142510")
	go responseSize("https://www.iana.org/help/example-domains")
	go responseSize("https://zhuanlan.zhihu.com")
	time.Sleep(time.Second * 5)
}
