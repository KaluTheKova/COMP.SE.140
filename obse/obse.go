package main

import (
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var allTopics string = "compse.*"
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"
var filename string = "/app/messages.txt"

// Subscribes to all messages within the network, therefore receiving from both compse140.o and compse140.i
// Stores the messages into a file
func main() {
	log.Printf("Observer starting. Sleeping 20 secs.")
	time.Sleep(20 * time.Second)

	consumeMessagesFromQueue()
}

func consumeMessagesFromQueue() {
	// initialize connection
	conn, err := amqp.Dial(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// open channel
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

	// Declare queue. In case consumer starts before publisher. We need to make sure queue exists.
	queue, err := ch.QueueDeclare(
		"mainQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Bind
	err = ch.QueueBind(
		"mainQueue",    // queue name
		"compse140.#",  // routing key
		"mainExchange", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	// Consume messages
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Printf("Listening to queue %s\n", queue.Name)

	var forever chan struct{}

	counter := 0

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s from queue %v", d.Body, queue.Name)
			counter++
			buildTimeStampedMessage(string(d.Body), counter)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Write listened messages to file
func writeToFile() {
	//file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0666)
	// 1. Create file if not exist
	// 2. Store file in container (separate mount/volume?)
	// 3. Append each received message to file

}

// Builds a message with timestamp and message counter
func buildTimeStampedMessage(message string, counter int) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z")
	timeStampedMessage := fmt.Sprintf("%v", timestamp)
	return timeStampedMessage
}

// Helper to check each ampq call
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Removes filename
func clearFileOnStartup(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Println(err)
	}
}
