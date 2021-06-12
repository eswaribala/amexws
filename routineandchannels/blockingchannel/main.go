package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://rediffmail.com",
		"http://rpsconsulting.in",
	}
	//create the channel
	c := make(chan string)
	for _, link := range links {
		//add go routine
		//pass channel as parameter
		go CheckLinks(link, c)
	}
	//receive data from channel
	fmt.Println(<-c)
	fmt.Println("Main Terminated....")
}

func CheckLinks(link string, c chan string) {
	//signal main that job done

	response, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "The site is down")
		//write to channel
		c <- "The Site is Down......"

	}

	go CheckUrl(link)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error Occurred in reading", err)
	}

	fmt.Println(len(body))
	//write to channel
	c <- "The Site is Up......"

}

func CheckUrl(url string) {
	fmt.Println("Received Url", url)
}
