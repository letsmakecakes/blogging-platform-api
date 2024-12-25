package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinLogrus is a middleware function that logs HTTP requests using Logrus
func GinLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer to measure request duration
		startTime := time.Now()

		// Process the request
		c.Next()

		// Calculate the duration of the request
		duration := time.Since(startTime)

		// Log the request details
		logger.WithFields(logrus.Fields{
			"status":    c.Writer.Status(),
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"ip":        c.ClientIP(),
			"latency":   duration.Milliseconds(),
			"userAgent": c.Request.UserAgent(),
			"error":     c.Errors.ByType(gin.ErrorTypePrivate).String(), // Logs any internal errors
		}).Info("HTTP request processed")
	}
}
