package handlers

import (
	"awesomeProject/config"
	"awesomeProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

// Accepts a request and prevents duplicate processing
func AcceptHandler(c *gin.Context) {
	id := c.Query("id")
	endpoint := c.Query("endpoint")

	if id == "" {
		c.String(http.StatusBadRequest, "failed")
		return
	}

	// Check if ID exists in Redis
	if config.Deduplication.Exists(config.Ctx, "id:"+id).Val() == 1 {
		c.String(http.StatusConflict, "request already exists \n")
	} else {
		// Store ID in Redis for 1 minute
		if _, err := config.Deduplication.SetNX(config.Ctx, "id:"+id, true, time.Minute).Result(); err != nil {
			c.String(http.StatusInternalServerError, "failed")
			return
		}
		c.String(http.StatusOK, "pushed succesfully \n")
	}

	//// Track unique requests
	config.RequestCount.Store(id, time.Now())
	c.String(http.StatusOK, "ok")

	// If an endpoint is provided, send count
	if endpoint != "" {
		resp := utils.SendCount(endpoint)
		if resp != nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading body:", err)
				return
			}
			fmt.Println("Response Data:", string(body))
			c.String(resp.StatusCode, string(body))
		}
	}
}
