package main

import (
	"fmt"
	"log"
	"net/http"
)

const SenseBaseURL = "https://api.sense.com/apiservice/api/v1"
const Port = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	// path := "/authenticate"
}

func main() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/", handler)
	log.Printf("Start server at http://localhost%s", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
