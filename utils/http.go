package utils

import (
	"log"
	"net/http"
	"strconv"

	"awesomeProject/config"
)

// Sends unique request count to the given endpoint
func SendCount(endpoint string) *http.Response {
	uniqueCount := 0
	config.RequestCount.Range(func(_, _ interface{}) bool {
		uniqueCount++
		return true
	})

	resp, err := http.Post(endpoint+"?count="+strconv.Itoa(uniqueCount), "application/json", nil)
	if err != nil {
		log.Println("Error sending count to endpoint:", err)
		return nil
	}

	log.Println("POST request sent to", endpoint, "Response Status:", resp.StatusCode)
	return resp
}
