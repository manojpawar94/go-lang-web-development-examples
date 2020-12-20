package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/schema"
)

//RenderTemplate handler
func renderIndexTemplate(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("public/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("Error occurred while executing the template or writing its output : ", err)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("public/login.html")
		parsedTemplate.Execute(w, nil)
	} else if r.Method == "POST" {
		user := readUserForm(r)
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}
}

func readUserForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	err := decoder.Decode(user, r.PostForm)
	if err != nil {
		log.Fatal("error mapping parsed form data to struct : ", err)
	}
	return user
}
