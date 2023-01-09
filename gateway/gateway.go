package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

const (
	filename        string = "messages.txt"
	httpservAddress string = "http://httpserv:8080"
	origAddress     string = "http://orig:8085"
)

// When requested, returns content of the file created by OBSE
func main() {
	log.Printf("API GATEWAY STARTING")

	router := gin.Default()
	router.GET("/messages", getMessages)
	router.PUT("/state", putState)
	router.GET("/state", getState)
	router.GET("/run-log", getRunLog)

	router.Run(":8083")
}

// GET /messages (as text/plain)
// Returns all message registered with OBSE-service. Assumed implementation
// forwards the request to HTTPSERV and returns the result.
func getMessages(ginContext *gin.Context) {
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
	// Read put payload
	payload, err := ioutil.ReadAll(ginContext.Request.Body)
	if err != nil {
		log.Panic(err)
	}

	payloadString := string(payload)

	log.Printf("Gateway payload: %s", payloadString)

	// Handle incorrect input
	if payloadString != "INIT" && payloadString != "PAUSED" && payloadString != "RUNNING" && payloadString != "SHUTDOWN" {
		ginContext.String(http.StatusOK, fmt.Sprintf("PUT %s not valid input", payloadString))
	}

	// Initialize client
	customClient := NewCustomClient()

	// Docker Client
	dockerClient := createDockerClient()
	defer dockerClient.Close()
	// GET CURRENT STATE. If current state == payload, nothing happens.
	//state := readState()

	// State caseswitch tänne
	// Cases
	// stateInit := "ORIG service set to initial state"
	// statePaused := "ORIG service paused"
	// stateRunning := "ORIG service running"
	// stateShutdown := "ORIG service shutting down"

	switch payloadString {
	case "INIT":
		// Clean message log, start origin from 0
		// Restart all containers from scratch?
		// https://gist.github.com/frikky/e2efcea6c733ea8d8d015b7fe8a91bf6
		resp := customClient.PutState(origAddress, string(payload))
		ginContext.String(http.StatusOK, resp)
	case "PAUSED":
		// Pause ORIG
		//resp := customClient.PutState(origAddress, string(payload))
		pauseContainer(dockerClient, "compse140-orig-1")
		ginContext.String(http.StatusOK, "ORIG paused\n")
	case "RUNNING":
		// Start ORIG
		//resp := customClient.PutState(origAddress, string(payload))
		unpauseContainer(dockerClient, "compse140-orig-1")
		ginContext.String(http.StatusOK, "ORIG running\n")
	case "SHUTDOWN":
		// Shutdown all containers
		// https://gist.github.com/frikky/e2efcea6c733ea8d8d015b7fe8a91bf6
		listContainers(dockerClient)
		ginContext.String(http.StatusOK, "Listed containers\n")
	}
}

// GET /state (as text/plain)
// get the value of state
func getState(c *gin.Context) {

}

// GET /run-log (as text/plain)
// Get information about state changes
func getRunLog(c *gin.Context) {

}

func readState(filename string) string {
	// file, err := os.OpenFile("state.txt")
	// if err != nil {
	// 	log.Panic(err)
	// }
	return "lol"
}

func readRunlog(filename string) {
	// Read runlog
}

// Write listened messages to file
func writeStateToFile(filename string, message string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		log.Panic(err)
	}

	// Flush writer
	file.Sync()

	//log.Printf("WROTE TO FILENAME %v MESSAGE %v\n", filename, message) // DEBUG
}

func createDockerClient() *client.Client {
	// client, err := client.NewEnvClient()
	// if err != nil {
	// 	fmt.Printf("Unable to create docker client: %s", err)
	// }

	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Panicf("Unable to create docker client: %s", err)
	}

	return client
}

func pauseContainer(client *client.Client, containerName string) {
	ctx := context.Background()

	err := client.ContainerPause(ctx, containerName)
	if err != nil {
		log.Panicf("Unable to stop container %v: %v", containerName, err)
	}
}

func unpauseContainer(client *client.Client, containerName string) {
	ctx := context.Background()

	err := client.ContainerUnpause(ctx, containerName)
	if err != nil {
		log.Panicf("Unable to stop container %s: %s", containerName, err)
	}
}

func listContainers(client *client.Client) {
	ctx := context.Background()
	container, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range container {
		fmt.Println(container.ID, container.Image)
	}
}
