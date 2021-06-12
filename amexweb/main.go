package main

import (
	"github.com/amexws/amexweb/models"
	"html/template"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7070"
)

func loadHome(w http.ResponseWriter, r *http.Request) {

	user := models.User{Id: 45832, Name: "Parameswari"}

	parsedTemplate, _ := template.ParseFiles("templates/first-template.html")
	err := parsedTemplate.Execute(w, user)
	if err != nil {
		log.Printf("Error occurred while executing the templateor writing its output : ", err)
		return
	}

	//fmt.Fprintf(w, "Launching Home Page!")
}
func main() {

	http.HandleFunc("/", loadHome)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Cannot load AMEX Home Page", err)
	}
}
