package middleware

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start timer
        startTime := time.Now()

        // Process request
        c.Next()

        // Calculate resolution time
        endTime := time.Now()
        latency := endTime.Sub(startTime)

        // Get request details
        method := c.Request.Method
        path := c.Request.URL.Path
        statusCode := c.Writer.Status()

        // Log request details
        log.Printf("Method: %s | Path: %s | Status: %d | Latency: %v", method, path, statusCode, latency)
    }
}
