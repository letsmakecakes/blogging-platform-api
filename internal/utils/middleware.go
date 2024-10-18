package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinLogrus is a middleware function that logs HTTP requests using Logrus
func GinLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		duration := time.Since(start)

		// Log fields
		logger.WithFields(logrus.Fields{
			"status":    c.Writer.Status(),
			"method":    c.Request.Method,
			"path":      c.Request.Method,
			"ip":        c.ClientIP(),
			"latency":   duration,
			"userAgent": c.Request.UserAgent(),
		}).Info("request processed")
	}
}
