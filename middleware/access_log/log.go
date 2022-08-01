package middleware

import (
	"rt-msg-carrier/log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerToFile() gin.HandlerFunc {
	// Instantiation
	logger := log.NewAccessLogger()
	return func(c *gin.Context) {
		//  Starting time
		startTime := time.Now()
		//  Processing requests
		c.Next()
		//  End time
		endTime := time.Now()
		//  execution time
		latencyTime := endTime.Sub(startTime)
		//  Request mode
		reqMethod := c.Request.Method
		//  Request routing
		reqUri := c.Request.RequestURI
		//  Status code
		statusCode := c.Writer.Status()
		//  request IP
		clientIP := c.ClientIP()
		//  Log format
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
