package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GLOBALS
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"
var routingKey = "compse140.o"
var i = 1

// Publishes messages to TOPIC compse140.o
// TOPIC compse140.o in RabbitMQ
func main() {
	log.Printf("ORIG STARTING") // DEBUG
	runService()
}

func runService() {
	log.Println("ORIG service running")
	conn, ch, err := initializeConnection(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
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
	if err != nil {
		runService()
	}

	// cancel when ended
	ctx, cancel := context.WithTimeout(context.Background(), 1800*time.Second)
	defer cancel()

	// Prefetch qos
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		runService()
	}

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
	if err != nil {
		runService()
	}
	log.Printf(" [x] Sent %s\n", body) // DEBUG
}

func initializeConnection(rabbitMQAddress string) (*amqp.Connection, *amqp.Channel, error) {
	var dialConfig amqp.Config
	dialConfig.Heartbeat = 10 * time.Second

	conn, err := amqp.DialConfig(rabbitMQAddress, dialConfig)
	if err != nil {
		return nil, nil, err
	}

	// create channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	return conn, ch, nil
}

func resetConnection(errMessage string) {
	log.Println(errMessage, ", resetting connection")
	conn, ch, err := initializeConnection(rabbitMQAddress)
	defer conn.Close()
	defer ch.Close()

	if err != nil {
		failOnError(err, "Resetting connection failed")
	}
}

// Helper to check each ampq call
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
