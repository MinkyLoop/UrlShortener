package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.Printf("%s %s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.RemoteAddr,
			time.Since(start))
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			}
		}()

		c.Next()
	}
}
