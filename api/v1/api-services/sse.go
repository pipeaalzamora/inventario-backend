package apiservices

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupSSe() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		// Critical headers for SSE to work with nginx
		gctx.Writer.Header().Set("Content-Type", "text/event-stream")
		gctx.Writer.Header().Set("Cache-Control", "no-cache")
		gctx.Writer.Header().Set("Connection", "keep-alive")
		gctx.Writer.Header().Set("Transfer-Encoding", "chunked")
		gctx.Writer.Header().Set("X-Accel-Buffering", "no") // Disable nginx buffering

		// CORS headers for SSE
		gctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		gctx.Writer.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

		// Force flush headers
		if f, ok := gctx.Writer.(http.Flusher); ok {
			f.Flush()
		}

		gctx.Next()
	}
}
