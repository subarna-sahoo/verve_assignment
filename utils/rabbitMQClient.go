package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

// InitRabbitMQ initializes the RabbitMQ connection once
func InitRabbitMQ() {
	// Connect to RabbitMQ with a retry mechanism
	var err error
	for i := 0; i < 5; i++ { // Retry up to 5 times
		RabbitMQConn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

		if err == nil {
			break
		}
		log.Printf("❌ RabbitMQ connection failed: %v. Retrying (%d/5)...", err, i+1)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	if err != nil {
		log.Printf("❌ Failed to connect to RabbitMQ after retries: %v", err)
		return
	}

	// Open a channel
	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Printf("❌ Failed to open RabbitMQ channel: %v", err)
		return
	}

	// Declare a queue
	_, err = RabbitMQChannel.QueueDeclare(
		"unique_requests", // queue name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Printf("❌ Failed to declare RabbitMQ queue: %v", err)
		return
	}

	fmt.Println("✅ RabbitMQ initialized successfully!")
}

// CloseRabbitMQ cleans up RabbitMQ resources
func CloseRabbitMQ() {
	if RabbitMQChannel != nil {
		RabbitMQChannel.Close()
	}
	if RabbitMQConn != nil {
		RabbitMQConn.Close()
	}
	log.Println("🛑 RabbitMQ connection closed")
}

// PublishToRabbitMQ sends the unique request count to the RabbitMQ queue
func PublishToRabbitMQ(count int) error {
	if RabbitMQChannel == nil {
		return fmt.Errorf("❌ RabbitMQ channel is not initialized")
	}

	message := fmt.Sprintf("Unique Requests in last minute: %d", count)
	err := RabbitMQChannel.Publish(
		"",                // exchange
		"unique_requests", // queue name (routing key)
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("❌ Failed to publish message to RabbitMQ: %w", err)
	}

	log.Printf("📩 Published message to RabbitMQ: %s", message)
	return nil
}
