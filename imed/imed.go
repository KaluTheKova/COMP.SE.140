package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var consumedRountingKey = "compse140.o"
var sendingRoutingKey = "compse140.i"
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"

// Subscribes for messages from compse140.o
// Publishes message to compse140.i
func main() {
	log.Printf("IMED starting.") // DEBUG

	consumeMessagesFromQueue()
}

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
		"",    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Bind
	err = ch.QueueBind(
		"",                  // queue name
		consumedRountingKey, // routing key
		"mainExchange",      // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	// // Prefect QoS
	// err = ch.Qos(
	// 	1,     // prefetch count
	// 	0,     // prefetch size
	// 	false, // global
	// )
	// failOnError(err, "Failed to set QoS")

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
			log.Printf("Received a message: %s from queue %v", d.Body, queue.Name) // DEBUG
			message := fmt.Sprintf("Got %v", string(d.Body))
			sendMessageToQueue(message)
			time.Sleep(1 * time.Second) // Wait for 1 second
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
		"",    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// cancel when ended
	ctx, cancel := context.WithTimeout(context.Background(), 1800*time.Second)
	defer cancel()

	// message body
	body := message
	err = ch.PublishWithContext(ctx,
		"mainExchange",    // exchange
		sendingRoutingKey, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s to mainQueue %v\n", body, queue.Name) // DEBUG
}

// Helper to check each ampq call
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
