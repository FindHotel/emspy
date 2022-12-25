package github

import (
	"log"

	"github.com/FindHotel/emspy/internal/app/server/handlers/interfaces"
	"github.com/gin-gonic/gin"
)

func WebhooksHandler(capturer interfaces.Capturer) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := c.GetRawData()
		if err != nil {
			log.Printf("Got error: %s", err)
		}
		capturer.Capture(c, r)
	}
}
