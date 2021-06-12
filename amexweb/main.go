package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7070"
)

func loadHome(w http.ResponseWriter, r *http.Request) {
	/*
	        result,_:=template.ParseFiles("/templates/index.html")
	        err:=result.Execute(w,nil);
	        if err!=nil{
	        	log.Fatal("Home Page cannot be loaded",err);
	   	 }
	*/
	fmt.Fprintf(w, "Launching Home Page!")
}
func main() {

	http.HandleFunc("/", loadHome)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Cannot load AMEX Home Page", err)
	}
}
