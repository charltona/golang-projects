package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range sites {
		go makeRequest(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			makeRequest(link, c)
		}(l)
	}

}

func makeRequest(link string, c chan string) {
	fmt.Println("checking", link)
	_, err := http.Get(link)

	if err != nil {
		fmt.Print("might be down! \n")
		c <- link
		return
	}

	fmt.Println(link, "status OK!")
	c <- link
}
