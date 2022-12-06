package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	filename        string = "messages.txt"
	httpservAddress string = "http://localhost:8080/"
)

type Client interface {
}

// When requested, returns content of the file created by OBSE
func main() {
	log.Printf("API GATEWAY STARTING")

	router := gin.Default()
	router.GET("/messages", getMessages)
	router.PUT("/state", putState)
	router.GET("/state", getState)
	router.GET("/run-log", getRunLog)

	router.Run("localhost:8083")
}

// GET /messages (as text/plain)
// Returns all message registered with OBSE-service. Assumed implementation
// forwards the request to HTTPSERV and returns the result.
func getMessages(ginContext *gin.Context) {
	log.Println("Received GET/messages") // DEBUG

	customClient := NewCustomClient()
	resp := customClient.GetMessages(httpservAddress)

	ginContext.String(http.StatusOK, string(resp))
}

// PUT /state (payload “INIT”, “PAUSED”, “RUNNING”, “SHUTDOWN”)
// PAUSED = ORIG service is not sending messages
// RUNNING = ORIG service sends messages
// If the new state is equal to previous nothing happens.
// There are two special cases:
// INIT = everything (except log information for /run-log and /messages) is in the
// initial state and ORIG starts sending again, state is set to RUNNING
// SHUTDOWN = all containers are stopped
func putState(ginContext *gin.Context) {
	log.Println("Received PUT/state") // DEBUG
	payload := ginContext.PostForm("LOL")

	log.Printf("Payload: %s", payload)

	// customClient := NewCustomClient()
	// resp := customClient.PutState(httpservAddress, "x")
	// ginContext.String(http.StatusOK, resp)
}

// GET /state (as text/plain)
// get the value of state
func getState(c *gin.Context) {

}

// GET /run-log (as text/plain)
// Get information about state changes
func getRunLog(c *gin.Context) {

}
