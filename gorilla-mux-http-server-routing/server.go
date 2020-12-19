package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//constants
const (
	ServerHost = "localhost"
	ServerPort = "8080"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/post", PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}",
		PathVariableHandler).Methods("GET", "PUT")
	router.Handle("/add/{num1}/{num2}", AddNumberHandler)
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
