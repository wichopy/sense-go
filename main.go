package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wichopy/sense-go/types"
)

const SenseBaseURL = "https://api.sense.com/apiservice/api/v1"
const Port = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	path := "/authenticate"

	resp, err := http.Post(SenseBaseURL+path, "application/x-www-form-urlencoded; charset=UTF-8;", r.Body)
	if err != nil {
		log.Fatalf("Error posting auth body, %v", err)
		return
	}

	log.Printf("Response: %v, %s, %s", resp.StatusCode, resp.Status, resp.Header)

	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body, %v", err)
		return
	}
	body, err := types.UnmarshalSenseAuthenticationResponse(rawBody)
	if err != nil {
		log.Fatalf("Error Unmarshalling response body, %v", err)
		return
	}

	log.Printf("Response:, %v", body)
}

func main() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/", handler)
	log.Printf("Start server at http://localhost%s", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
