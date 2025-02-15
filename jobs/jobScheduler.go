package jobs

import (
	"log"
	"time"
	"verve_assignment/models"
	"verve_assignment/utils"
)

// StartJobScheduler runs a job at a specified interval
func StartJobScheduler(interval time.Duration, job func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		job()
	}
}

// UniqueRequestJob fetches unique request IDs from Redis and publishes count to RabbitMQ
func UniqueRequestJob() {
	ids, err := models.GetUniqueRequestIDsFromLastMinute()
	if err != nil {
		log.Println("Error fetching unique requests:", err)
		return
	}

	count := len(ids)

	// Publish count to RabbitMQ
	err = utils.PublishToRabbitMQ(count)
	if err != nil {
		log.Println("❌ Failed to publish message to RabbitMQ:", err)
	} else {
		log.Printf("✅ Successfully sent unique request count to RabbitMQ: %d\n", count)
	}
}
