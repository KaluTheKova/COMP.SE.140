package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

// GLOBALS
var currentState string = "INIT"

// CONST
const (
	filename        string = "state.txt"
	httpservAddress string = "http://httpserv:8080"
	origAddress     string = "http://orig:8085"
)

// When requested, returns content of the file created by OBSE
func main() {
	log.Printf("API GATEWAY STARTING")

	clearFileOnStartup(filename)

	// Write initial state
	writeStateToFile(filename, "INIT")

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

	// Initialize HTTP client
	//customClient := NewCustomClient()

	// Docker Client
	dockerClient := createDockerClient()
	defer dockerClient.Close()

	// GET CURRENT STATE. If current state == payload, nothing happens.
	if payloadString == currentState {
		ginContext.String(http.StatusOK, fmt.Sprintf("Current state already %v", payloadString))
		log.Printf("Current state already %v\n", payloadString)
		return
	} else if payloadString == "INIT" && currentState != "INIT" { // INIT is special case according to instructions. INIT should set state to RUNNING
		writeStateToFile(filename, "RUNNING")
		currentState = "RUNNING"
	} else {
		writeStateToFile(filename, payloadString)
		currentState = payloadString
	}

	// Cases
	stateInit := "ORIG service set to initial state\n"
	statePaused := "ORIG service paused\n"
	stateRunning := "ORIG service running\n"
	stateShutdown := "ORIG service shutting down\n"

	switch payloadString {
	case "INIT":
		// Restart container ORIG
		restartContainers(dockerClient, "compse140-orig-1")
		//resp := customClient.PutState(origAddress, string(payload))
		ginContext.String(http.StatusOK, stateInit)
	case "PAUSED":
		// Pause ORIG
		pauseContainer(dockerClient, "compse140-orig-1")
		ginContext.String(http.StatusOK, statePaused)
	case "RUNNING":
		// Restart ORIG
		unpauseContainer(dockerClient, "compse140-orig-1")
		ginContext.String(http.StatusOK, stateRunning)
	case "SHUTDOWN":
		// Shutdown all containers
		ginContext.String(http.StatusOK, stateShutdown)
		stopAllContainers(dockerClient)
	}
}

// GET /state (as text/plain)
// get the value of state
func getState(ginContext *gin.Context) {
	ginContext.String(http.StatusOK, currentState+"\n")
}

// GET /run-log (as text/plain)
// Get information about state changes
func getRunLog(ginContext *gin.Context) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	ginContext.String(http.StatusOK, string(fileContents))
}

// Write state to file
func writeStateToFile(filename string, state string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	timeStampedState := buildTimeStampedState(state)

	_, err = file.WriteString(timeStampedState + "\n")
	if err != nil {
		log.Panic(err)
	}

	// Flush writer
	file.Sync()
}

func createDockerClient() *client.Client {
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

func restartContainers(client *client.Client, containerName string) {
	ctx := context.Background()

	timeout := 0
	options := container.StopOptions{Signal: "SIGTERM", Timeout: &timeout}
	//options := container.StopOptions{"SIGTERM", &timeout}

	err := client.ContainerRestart(ctx, containerName, options)
	if err != nil {
		panic(err)
	}
}

// Removes filename
func clearFileOnStartup(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Removed file: %v", filename) // DEBUG
}

// Builds a message with timestamp and message counter
func buildTimeStampedState(state string) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z")
	timeStampedMessage := fmt.Sprintf("%v %v ", timestamp, state)
	return timeStampedMessage
}

func stopAllContainers(client *client.Client) {
	ctx := context.Background()

	containers, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	timeout := 0
	options := container.StopOptions{Timeout: &timeout}

	// Shutdown all containers whose image name contains substring "compse140", but shutdown Gateway only after everything else is stopped.
	for _, container := range containers {
		if strings.Contains(container.Image, "compse140") && !strings.Contains(container.Image, "compse140-gateway") {
			log.Print("Stopping container ", container.Image, "... ") // DEBUG
			//ginContext.String(http.StatusOK, "Stopping container ", container.Image, "... ")
			err := client.ContainerStop(ctx, container.ID, options)
			if err != nil {
				panic(err)
			}
			log.Println("Success") // DEBUG
			//ginContext.String(http.StatusOK, "Success")
		}
	}
	containers, err = client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	// Last, shutdown Gateway
	for _, container := range containers {
		if strings.Contains(container.Image, "compse140-gateway") {
			log.Print("Stopping container ", container.Image, "... ") // DEBUG
			//ginContext.String(http.StatusOK, "Stopping container ", container.Image, "... ")
			err := client.ContainerStop(ctx, container.ID, options)
			if err != nil {
				panic(err)
			}
			log.Println("Success") // DEBUG
			//ginContext.String(http.StatusOK, "Success")
		}
	}
}

func listContainers(client *client.Client) {
	ctx := context.Background()
	container, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range container {
		log.Println(container.ID, container.Image)
	}
}
