package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

//const for Go HTTP Server
const (
	HOST           = "localhost"
	PORT           = "8080"
	AdminUser      = "admin"
	AdminPassword  = "admin"
	BasicAuthRealm = "Please enter your username and password"
)

func main() {
	http.HandleFunc("/", BasicAuth(HelloWorldHandler, BasicAuthRealm))
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("Error while starting the http server: ", err)
	}
}

//HelloWorldHandler function to handle the Hello World Request
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World it is go lang program")
}

//BasicAuth act as filter to authenticate the user access to application
func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(AdminUser)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(AdminPassword)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler(w, r)
	}
}
