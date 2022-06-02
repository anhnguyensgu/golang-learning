package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeNetworkCall(url string, result chan string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error when requesting url " + url)
	}

	fmt.Printf("calling %s \n", url)
	body, ioErr := ioutil.ReadAll(response.Body)
	if ioErr != nil {
		log.Fatal("cannot parse body")
	}

	fmt.Printf("body size %d\n", len(body))
	result <- url
}

func main() {
	links := []string{"https://stackoverflow.com", "https://google.com.vn"}
	chans := make(chan string, 10)
	for _, link := range links {
		go makeNetworkCall(link, chans)
	}

	for i := 0; i < len(links); i++ {
		fmt.Printf("done with link %s\n", <-chans)
	}
}
