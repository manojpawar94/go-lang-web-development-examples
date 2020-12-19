package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//constants
const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

func main() {
	router := mux.NewRouter()

	//console logs
	router.Handle("/", handlers.LoggingHandler(os.Stdout,
		http.HandlerFunc(GetRequestHandler))).Methods("GET")

	//file logs
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}

	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}",
		handlers.LoggingHandler(logFile, PathVariableHandler)).Methods("GET", "PUT")
	router.Handle("/add/{num1}/{num2}", handlers.LoggingHandler(logFile, AddNumberHandler))

	http.ListenAndServe(ServerHost+":"+ServerPort, router)
}

//GetRequestHandler handler
var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

//PostRequestHandler handler
var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a Post Request!"))
	})

//PathVariableHandler handler
var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		w.Write([]byte("Hi " + name))
	})

//AddNumberHandler handler
var AddNumberHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		num1, err := strconv.Atoi(mux.Vars(r)["num1"])
		num2, err := strconv.Atoi(mux.Vars(r)["num2"])

		if err != nil {
			w.Write([]byte("Invalid request " + err.Error()))
		}
		result := num1 + num2
		w.Write([]byte("Result: " + strconv.Itoa(result)))
	})
