package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//constants
const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", renderIndexTemplate).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	err := http.ListenAndServe(ServerHost+":"+ServerPort, router)

	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
