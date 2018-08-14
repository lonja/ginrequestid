//
// ginrequestid
//
// Set an UUID4 string as Request ID into response headers ("X-Request-Id") and
// expose that value as "RequestId" in order to use it inside the application for logging
// or propagation to other systems.
//
package ginrequestid

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"math/rand"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with random hex string
		if requestID == "" {
			requestID = randomStr(8)
		}

		// Expose it for use in the application
		c.Set("request_id", requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}

func randomStr(len int) string {
	buffer := make([]byte, len)
	rand.Read(buffer)
	return fmt.Sprintf("%x", buffer)
}
