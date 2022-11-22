package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	filename        string = "messages.txt"
	httpservAddress string = "http://localhost:8080/"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
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
func getMessages(c *gin.Context) {
	log.Println("Received GET/messages") // DEBUG
	resp, err := http.Get(httpservAddress)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	log.Println("Reading response") // DEBUG
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	c.String(http.StatusOK, string(respBody))
}

// GetMessagesAPICall executes http.Get to given address and returns bytecontent
/* func getMessagesFromHttpserv(url string) []byte {

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
} */

// PUT /state (payload “INIT”, “PAUSED”, “RUNNING”, “SHUTDOWN”)
// PAUSED = ORIG service is not sending messages
// RUNNING = ORIG service sends messages
// If the new state is equal to previous nothing happens.
// There are two special cases:
// INIT = everything (except log information for /run-log and /messages) is in the
// initial state and ORIG starts sending again, state is set to RUNNING
// SHUTDOWN = all containers are stopped
func putState(c *gin.Context) {

}

// GET /state (as text/plain)
// get the value of state
func getState(c *gin.Context) {

}

// GET /run-log (as text/plain)
// Get information about state changes
func getRunLog(c *gin.Context) {

}
