package handlers

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Returns the count of unique requests
func GetHandler(c *gin.Context) {
	countFromRedis := len(config.Deduplication.Keys(config.Ctx, "id:*").Val())
	c.JSON(http.StatusOK, gin.H{"Request count ": countFromRedis})
}
