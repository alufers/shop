package shop

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func errorPrinterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors
		if len(errs) > 0 {
			data, _ := json.Marshal(map[string]interface{}{
				"Errors": errs.Errors(),
			})
			c.Writer.Write(data)
		}
	}
}
