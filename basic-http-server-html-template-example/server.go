package main

import (
	"html/template"
	"log"
	"net/http"
)

//constants
const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

//RenderTemplate handler
func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{ID: "1", Name: "Manoj Pawar"}
	parsedTemplate, _ := template.ParseFiles("public/index.html")

	err := parsedTemplate.Execute(w, person)

	if err != nil {
		log.Fatal("Error occurred while executing the template or writing its output : ", err)
		return
	}
}

func main() {

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", RenderTemplate)
	err := http.ListenAndServe(ServerHost+":"+ServerPort, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
