package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRequestID(c *gin.Context) string {
	if id := c.GetString("request_id"); id != "" {
		return id
	}
	return uuid.NewString()
}
