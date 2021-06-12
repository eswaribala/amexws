package main

import (
	"fmt"
	"io/ioutil"
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
	//add counter size
	wg.Add(len(links))
	for _, link := range links {
		//add go routine
		go CheckLinks(link)
	}
	//make main to wait
	wg.Wait()
	fmt.Println("Main Terminated....")
}

func CheckLinks(link string) {
	//signal main that job done
	defer wg.Done()
	response, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "The site is down")

	}

	go CheckUrl(link)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error Occurred in reading", err)
	}

	fmt.Println(len(body))

}

func CheckUrl(url string) {
	fmt.Println("Received Url", url)
}
