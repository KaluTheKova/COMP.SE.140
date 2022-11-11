package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var allTopics string = "compse.*"
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"
var filename string = "messages.txt"
var path string = "/app"

// Subscribes to all messages within the network, therefore receiving from both compse140.o and compse140.i
// Stores the messages into a file
func main() {
	log.Printf("Observer starting.") // DEBUG

	clearFileOnStartup("messages.txt")

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

	log.Printf("Listening to queue %s\n", queue.Name) // DEBUG

	var forever chan struct{}

	counter := 0

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s from queue %v", d.Body, queue.Name) // DEBUG
			counter++
			timeStampedMessage := buildTimeStampedMessage(string(d.Body), counter, d.RoutingKey)
			err := writeToFile(filename, timeStampedMessage)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Write listened messages to file
func writeToFile(filename string, message string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		return err
	}

	// Flush writer
	file.Sync()

	log.Printf("WROTE TO FILENAME %v MESSAGE %v\n", filename, message) // DEBUG

	return nil
}

// Builds a message with timestamp and message counter
func buildTimeStampedMessage(message string, counter int, topic string) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z")
	timeStampedMessage := fmt.Sprintf("%v %v %v to %v", timestamp, counter, message, topic)
	return timeStampedMessage
}

// Removes filename
func clearFileOnStartup(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Removed file: %v", filename) // DEBUG
}

// Helper to check each ampq call
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func listAllFilesInDirectory(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
