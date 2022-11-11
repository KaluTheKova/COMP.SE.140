package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var consumedQueue = "compse140.o"
var sendingQueue = "compse140.i"
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"

// Subscribes for messages from compse140.o
// Publishes message to compse140.i
func main() {
	log.Printf("IMED starting. Sleeping 20 secs.")
	time.Sleep(20 * time.Second)

	consumeMessagesFromQueue()
}

// Consumes messages from compse140.o
// and sends them to compse140.i
func consumeMessagesFromQueue() {
	// initialize connection
	conn, err := amqp.Dial(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Separate channels for consume and publish
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
		consumedQueue,  // routing key
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

	go func() {
		for d := range msgs {
			time.Sleep(1 * time.Second) // Wait for 1 second
			log.Printf("Received a message: %s from queue %v", d.Body, queue.Name)
			message := fmt.Sprintf("Got %v", string(d.Body))
			sendMessageToQueue(message)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func sendMessageToQueue(message string) {
	// initialize connection
	conn, err := amqp.Dial(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Separate channels for consume and publish
	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// declare queue
	queue, err := ch.QueueDeclare(
		"mainQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// cancel when ended
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
	log.Printf(" [x] Sent %s to topic %v\n", body, queue.Name) // DEBUG
}

// Helper to check each ampq call
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
