package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"mikou/global"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().UnixNano()
		c.Next()
		endTime := time.Now().UnixNano()

		//fields := logger.Fields{
		//	"request":  c.Request.PostForm.Encode(),
		//	"response": bodyWriter.body.String(),
		//}
		sepend := float64(endTime - beginTime) / 1e6
		global.LoggerV2.With("request: ", c.Request.PostForm.Encode(), "response: ", bodyWriter.body.String(), ).Infof("URL: %s  spend: %f ms",
			c.Request.URL,
			sepend,
		)
		//global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
		//	c.Request.Method,
		//	bodyWriter.Status(),
		//	beginTime,
		//	endTime,
		//)
	}
}
