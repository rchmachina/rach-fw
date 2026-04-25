package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rchmachina/rach-fw/internal/utils/helper"
)

///// this helper only used for handler api dont use for other purpose like grpc or other delivery method

type Meta struct {
	Page      int `json:"page,omitempty"`
	Limit     int `json:"limit,omitempty"`
	TotalData int `json:"total_data,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}

type APIResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Meta      *Meta       `json:"meta,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	RequestID string      `json:"requestId"`
	Timestamp int64       `json:"timestamp"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}, meta *Meta) {
	c.JSON(http.StatusOK, APIResponse{
		Message:   message,
		Data:      data,
		Meta:      meta,
		RequestID: helper.GetRequestID(c.Request.Context()),
		Timestamp: time.Now().Unix(),
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, APIResponse{
		Message:   message,
		Error:     err.Error(),
		RequestID: helper.GetRequestID(c.Request.Context()),
		Timestamp: time.Now().Unix(),
	})
}
