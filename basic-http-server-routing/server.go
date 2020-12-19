package main

import (
	"fmt"
	"log"
	"net/http"
)

//constants
const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	err := http.ListenAndServe(ServerHost+":"+ServerPort, nil)

	if err != nil {
		log.Fatal("Error starting http server ", err)
	}
}

//helloWorld
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

//login
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
}

//logout
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Page")
}
