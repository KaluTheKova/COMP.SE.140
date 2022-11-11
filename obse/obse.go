package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// TO DO:
// Käy tutoriaali läpi. Viestit pitää mennä exchangeen. Sitten saat tän homman toimimaan :)

// GLOBALS
var topicO = "compse140.o"
var topicI = "compse140.i"
var rabbitMQAddress string = "amqp://guest:guest@rabbitmq:5672/"

// Subscribes for all messages within the network, therefore receiving from both compse140.o and compse140.i
// Stores the messages into a file
func main() {
	log.Printf("Observer starting. Sleeping 20 secs.")
	time.Sleep(20 * time.Second)

	// TO DO: TOPIC compse140.o and TOPIC compse140.i
	consumeMessagesFromQueue2("compse140.i") // Ongelma siis täällä. Obse ei pääse kuuntelemaan topicia "compse140.i"
	consumeMessagesFromQueue1("compse140.o") // Jälkimmäinen consumer ei ikinä käynnisty
}

func consumeMessagesFromQueue1(queueName string) {
	log.Println("DEBUG: STARTING TOPIC2 CONSUMER")
	// initialize connection
	conn, err := amqp.Dial(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Separate channels for consume and publish
	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare queue. In case consumer starts before publisher. We need to make sure queue exists.
	queue, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Prefetch
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// Consume messages
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack OFF, send MANUAL ack in worker
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			//time.Sleep(1 * time.Second) // Wait for 1 second
			log.Printf("Received a message: %s from topic %v", d.Body, queue.Name)
			d.Ack(false) // ACKNOWLEDGE
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func consumeMessagesFromQueue2(queueName string) {
	log.Println("DEBUG: STARTING TOPIC2 CONSUMER")

	// initialize connection
	conn, err := amqp.Dial(rabbitMQAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Separate channels for consume and publish
	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare queue. In case consumer starts before publisher. We need to make sure queue exists.
	queue, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Prefetch
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// Consume messages
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Printf("Listening to topic %v\n", queue.Name)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			//time.Sleep(1 * time.Second) // Wait for 1 second
			log.Printf("Received a message: %s from topic %v", d.Body, queue.Name)
			d.Ack(false) // ACKNOWLEDGE
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Write listened messages to file
func writeToFile() {
	// 1. Create file if not exist
	// 2. Store file in container (separate mount/volume?)
	// 3. Append each received message to file

}

func buildTimeStampedMessage(message string) string {
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
