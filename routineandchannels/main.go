package main

import (
	"fmt"
	"net/http"
	"sync"
)

//create wait group
var wg sync.WaitGroup

func main() {

	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://rediffmail.com",
		"http://rpsconsulting.in",
	}

	wg.Add(len(links))
	for _, link := range links {
		//add go routine
		go CheckLinks(link)
	}
	wg.Wait()
	fmt.Println("Main Terminated....")
}

func CheckLinks(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "The site is down")

	} else {
		fmt.Println(link, "The site is up")
	}

}
