package main

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	lmt := tollbooth.NewLimiter(2, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second * 5})
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	lmt.SetMethods([]string{"POST", "GET"}) //放开更精准限制，但是也放松了流量。

	r.Use(LimitHandler(lmt))
	r.GET("/index", func(c *gin.Context) {
		c.String(200, "Get Hello, world!")
	})
	r.POST("/index", func(c *gin.Context) {
		c.String(200, "Post Hello, world!")
	})
	r.Run(":12345")
}

func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}