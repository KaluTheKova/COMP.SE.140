package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
// var rabbitMQAddress string = os.Getenv("rabbitMQAddress")
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"
var sendingQueue = "compse140.o"

// Publishes messages to TOPIC compse140.o
// TOPIC compse140.o in RabbitMQ
func main() {

	log.Printf("Original starting. Sleeping 20 secs.")
	time.Sleep(20 * time.Second)

	conn := initializeConnection(rabbitMQAddress)
	defer conn.Close()

	// Send 3 messages to rabbitMQ
	numOfMessages := 3
	for i := 1; i < numOfMessages+1; i++ {
		message := createMessages(i)
		sendMessageToRabbit(message, sendingQueue, conn)
		time.Sleep(3 * time.Second)
	}
}

// createMessages Creates and returns string "MSG_{%v}" where %v is the int given as parameter
func createMessages(numOfMessage int) string {
	message := fmt.Sprintf("MSG_{%v}", numOfMessage)
	return message
}

func sendMessageToRabbit(message string, sendingQueue string, conn *amqp.Connection) {
	// create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// declare queue
	queue, err := ch.QueueDeclare(
		sendingQueue, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// cancel when ended
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// message body
	body := message
	err = ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
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
