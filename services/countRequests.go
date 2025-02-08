package services

import (
	"log"
	"sync"
	"time"

	"awesomeProject/config"
)

// Logs the unique request count every minute
func LogUniqueCounts() {
	for {
		time.Sleep(time.Minute)
		uniqueCount := 0
		config.RequestCount.Range(func(_, _ interface{}) bool {
			uniqueCount++
			return true
		})

		log.SetOutput(config.LogFile)
		log.Println("Unique Requests in last minute:", uniqueCount)

		// Reset the counter
		config.RequestCount = sync.Map{}
	}
}
