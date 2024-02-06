package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		// Set example variable
		c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("latency: %s", latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Printf("status: %s", strconv.Itoa(status))
		log.Printf("response body: %s", blw.body.String())
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
