package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/icholy/digest"
	"github.com/rs/cors"
)

var (
	url      = ""
	username = ""
	password = ""
)

/*
Ref: https://www.digital-loggers.com/rest.html
*/

/*
curl --digest -u admin:1234 -X PUT -H "X-CSRF: x" --data "value=true" "http://192.168.0.100/restapi/relay/outlets/all;/state/"
*/
func allRelaysSetTo(value bool) {
	path := "/restapi/relay/outlets/all;/state/"
	data := fmt.Sprintf("value=%v", value)

	// Create a new request
	req, err := http.NewRequest("PUT", url+path, strings.NewReader(data))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers
	req.Header.Set("X-CSRF", "x")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a Digest Transport
	transport := &digest.Transport{
		Username: username,
		Password: password,
	}

	// Create an HTTP client with the Digest Transport
	client := &http.Client{
		Transport: transport,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Printf("Response status: %s\n", resp.Status)
}

/*
curl -u admin:1234 -H "Accept:application/json" --digest "http://192.168.0.100/restapi/relay/outlets/all;/physical_state/"
*/
func getAllRelaysPhysicalState() []bool {
	path := "/restapi/relay/outlets/all;/physical_state/"

	// Create a new request
	req, err := http.NewRequest("GET", url+path, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil
	}

	// Set headers
	req.Header.Set("Accept", "application/json")

	// Create a Digest Transport
	transport := &digest.Transport{
		Username: username,
		Password: password,
	}

	// Create an HTTP client with the Digest Transport
	client := &http.Client{
		Transport: transport,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// Parse and print the response body as JSON
	var physicalState []bool
	err = json.NewDecoder(resp.Body).Decode(&physicalState)
	if err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		return nil
	}

	fmt.Printf("Response status: %s\n", resp.Status)
	fmt.Printf("Physical state: %v\n", physicalState)

	return physicalState
}

func relayStateHandler(w http.ResponseWriter, r *http.Request) {
	state := getAllRelaysPhysicalState()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(state)
}

func main() {
	allRelaysSetTo(true)
	getAllRelaysPhysicalState()
	time.Sleep(time.Minute)
	allRelaysSetTo(false)
	getAllRelaysPhysicalState()

	mux := http.NewServeMux()
	mux.HandleFunc("/relays/state", relayStateHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(mux)

	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
