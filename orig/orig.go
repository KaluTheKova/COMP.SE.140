package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"
var sendingQueue = "compse140.o"

// Publishes messages to TOPIC compse140.o
// TOPIC compse140.o in RabbitMQ
func main() {
	log.Printf("ORIG STARTING") // DEBUG
	router := gin.Default()
	router.PUT("/", stateHandler)
	router.Run(":8085")
}

// Act based on PUT request payload
func stateHandler(ginContext *gin.Context) {
	log.Println("ORIG received PUT/state") // DEBUG

	// Read put payload
	payload, err := ioutil.ReadAll(ginContext.Request.Body)
	if err != nil {
		log.Panic(err)
	}

	payloadString := string(payload)

	log.Printf("ORIG payload: %s", payloadString)

	// Cases
	statePaused := "ORIG service paused"
	stateRunning := "ORIG service running"

	switch payloadString {
	case "PAUSED":
		ginContext.String(http.StatusOK, statePaused)
		pauseService()
	case "RUNNING":
		ginContext.String(http.StatusOK, stateRunning)
		runService()
	}
}

func pauseService() {
	log.Println("ORIG service paused") // DEBUG
}

func runService() {
	log.Println("ORIG service running")
	conn := initializeConnection(rabbitMQAddress)
	defer conn.Close()

	// Send messages forever TO DO
	sendMessagesForever := true
	i := 1
	for sendMessagesForever == true {
		message := createMessages(i)
		sendMessageToRabbit(message, conn)
		time.Sleep(3 * time.Second) // wait 3 seconds
		i++
	}
}

// createMessages Creates and returns string "MSG_{%v}" where %v is the int given as parameter
func createMessages(numOfMessage int) string {
	message := fmt.Sprintf("MSG_{%v}", numOfMessage)
	return message
}

func sendMessageToRabbit(message string, conn *amqp.Connection) {
	// create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Exchange
	err = ch.ExchangeDeclare(
		"mainExchange", // name
		"topic",        // type TOPIC?
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// cancel when ended
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prefetch qos
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// message body
	body := message
	err = ch.PublishWithContext(ctx,
		"mainExchange", // exchange
		sendingQueue,   // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body) // DEBUG
}

func initializeConnection(rabbitMQAddress string) *amqp.Connection {
	conn, err := amqp.Dial(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	return conn
}

// Helper to check each ampq call
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
