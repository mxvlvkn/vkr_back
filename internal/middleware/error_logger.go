package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()

		if status < 400 && len(c.Errors) == 0 {
			log.Printf("SUCCESS | %s | %3d | %13v | %s %s",
				ip, status, duration, method, path)
			return
		}

		log.Printf("ERROR | %s | %3d | %13v | %s %s",
			ip, status, duration, method, path)

		for _, err := range c.Errors {
			log.Printf(" → %v;", err.Err)
		}
	}
}