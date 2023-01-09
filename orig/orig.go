package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"
var routingKey = "compse140.o"
var conn, connErr = amqp.Dial(rabbitMQAddress)
var ch, chErr = conn.Channel()
var start = make(chan struct{})
var pause = make(chan struct{})
var quit = make(chan struct{})
var wg sync.WaitGroup
var i = 1

// Publishes messages to TOPIC compse140.o
// TOPIC compse140.o in RabbitMQ
func main() {
	log.Printf("ORIG STARTING") // DEBUG
	runService()
	// router := gin.Default()
	// router.PUT("/", stateHandler)
	// router.Run(":8085")
}

// // Act based on PUT request payload
// func stateHandler(ginContext *gin.Context) {
// 	runService()

// 	// Read put payload
// 	payload, err := ioutil.ReadAll(ginContext.Request.Body)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	payloadString := string(payload)

// 	log.Printf("ORIG payload: %s", payloadString)

// 	// Cases
// 	statePaused := "ORIG service paused\n"
// 	stateRunning := "ORIG service running\n"

// }

// func routine() {
// 	for {
// 		select {
// 		case <-pause:
// 			log.Println("ORIG service paused")
// 			// ch.Cancel("", true)
// 			// conn.Close()
// 			select {
// 			case <-play:
// 				log.Println("ORIG service running")
// 				runService()
// 			case <-quit:
// 				wg.Done()
// 				return
// 			}
// 		case <-quit:
// 			wg.Done()
// 			return
// 		default:
// 			runService()
// 		}
// 	}
// }

// func pauseService() {
// 	ch.Cancel()
// 	log.Println("ORIG service paused") // DEBUG
// 	conn.Close()
// 	ch.Close()
// 	return
// }

func runService() {
	log.Println("ORIG service running")
	conn := initializeConnection(rabbitMQAddress)
	defer conn.Close()

	// create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	for {
		message := createMessages(i)
		sendMessageToRabbit(message, ch)
		time.Sleep(3 * time.Second) // wait 3 seconds
		i++
	}

}

// createMessages Creates and returns string "MSG_{%v}" where %v is the int given as parameter
func createMessages(numOfMessage int) string {
	message := fmt.Sprintf("MSG_{%v}", numOfMessage)
	return message
}

func sendMessageToRabbit(message string, ch *amqp.Channel) {
	// // create channel
	// ch, err := conn.Channel()
	// failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	// Exchange
	err := ch.ExchangeDeclare(
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
		routingKey,     // routing key
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
