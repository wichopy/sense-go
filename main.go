package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/wichopy/sense-go/app"
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

	// log.Printf("Response: %v, %s, %s", resp.StatusCode, resp.Status, resp.Header)

	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body, %v", err)
		return
	}
	body, err := app.UnmarshalSenseAuthenticationResponse(rawBody)
	if err != nil {
		log.Fatalf("Error Unmarshalling response body, %v", err)
		return
	}

	auth := app.Auth{
		AccessToken: body.AccessToken,
		AccountID:   body.AccountID,
		UserID:      body.UserID,
		// FIXME: Hard coded getting the first monitor.
		MonitorID: body.Monitors[0].ID,
	}
	// log.Printf("Response:, %v", body)
	// log.Printf("Auth:, %v", auth)

	realtimeClient(auth.MonitorID, auth.AccessToken)
}

func realtimeClient(monitorId int64, accessToken string) {
	addr := "clientrt.sense.com"
	path := fmt.Sprintf("monitors/%s/realtimefeed", strconv.FormatInt(monitorId, 10))
	query := fmt.Sprintf("access_token=%s", accessToken)
	u := url.URL{Scheme: "wss", Host: addr, Path: path, RawQuery: query}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	log.Printf("Connected to web socket, waiting for specified devices to toggle...")
	// defer c.Close()

	// done := make(chan struct{})
	// go func() {
	// defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		m, err := app.UnmarshalSenseRealTimeMessage(message)

		for _, device := range m.Payload.Devices {
			// FIXME: Use a configurable list of devices to watch.
			if device.Name == "Fridge" {
				log.Printf("Fridge is on")
			}
		}
	}
	// }()

}

func main() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/", handler)
	log.Printf("Start server at http://localhost%s", Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
