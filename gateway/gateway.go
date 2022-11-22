package main

import (
	"io"
	"log"
	"net/http"
)

const (
	filename        string = "messages.txt"
	httpservAddress string = "http://localhost:8080/"
)

// When requested, returns content of the file created by OBSE
func main() {
	log.Printf("API GATEWAY STARTING")
	http.HandleFunc("/messages", getMessages)
	//http.HandleFunc("/state", putState)
	//http.HandleFunc("/state", getState)
	//http.HandleFunc("/run-log", getRunLog)
	http.ListenAndServe(":8083", nil)
}

// GET /messages (as text/plain)
// Returns all message registered with OBSE-service. Assumed implementation
// forwards the request to HTTPSERV and returns the result.
func getMessages(writer http.ResponseWriter, req *http.Request) {
	respBody := GetMessagesAPICall(httpservAddress)

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/text")
	writer.Write([]byte(respBody))
	return
}

// GetMessagesAPICall sends http.Get to given address and returns bytecontent
func GetMessagesAPICall(url string) []byte {

	log.Println("Received GET/messages") // DEBUG
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	log.Println("Reading response") // DEBUG
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return respBody
}

// PUT /state (payload “INIT”, “PAUSED”, “RUNNING”, “SHUTDOWN”)
// PAUSED = ORIG service is not sending messages
// RUNNING = ORIG service sends messages
// If the new state is equal to previous nothing happens.
// There are two special cases:
// INIT = everything (except log information for /run-log and /messages) is in the
// initial state and ORIG starts sending again,
//  state is set to RUNNING
// SHUTDOWN = all containers are stopped

// GET /state (as text/plain)
// get the value of state

// GET /run-log (as text/plain)
// Get information about state changes
