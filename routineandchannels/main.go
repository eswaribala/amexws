package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://rediffmail.com",
		"http://rpsconsulting.in",
	}

	for _, link := range links {
		CheckLinks(link)
	}

}

func CheckLinks(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "The site is down")

	} else {
		fmt.Println(link, "The site is up")
	}

}
