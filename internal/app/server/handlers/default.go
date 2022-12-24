package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterDefaultHandlers(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})

	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy")
	})
}
