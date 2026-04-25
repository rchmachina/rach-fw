package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	constanta "github.com/rchmachina/rach-fw/internal/const"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	log "github.com/rchmachina/rach-fw/internal/infrastructure/logger"
)

func IncomingRequest(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		info := fmt.Sprintf("[HTTP] URL: %s \n METHOD: %s DURATION: %s", c.Request.URL.Path, c.Request.Method, duration.String())

		logger := logger.WithCtx(c.Request.Context())
		toField := log.ToField("IncomingRequest", info)
		logger.Info("IncomingRequest", toField)

		println()
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.NewString()
		c.Set("request_id", reqID)
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, constanta.RequestIDKey, reqID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Request-ID", reqID)
		c.Next()
	}
}
