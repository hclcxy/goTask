package middleware

import (
	"go-task4/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		logger.NewGormLogger().Info(
			c.Request.Context(),
			"request handled",
			"method", method,
			"path", path,
			"status", status,
			"latency", latency,
		)
	}
}
