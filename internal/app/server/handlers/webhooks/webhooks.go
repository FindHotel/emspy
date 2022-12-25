package webhooks

import (
	"github.com/FindHotel/emspy/internal/app/server/handlers/webhooks/github"
	"github.com/FindHotel/emspy/internal/app/server/handlers/webhooks/shortcut"
	"github.com/FindHotel/emspy/internal/app/store"
	"github.com/gin-gonic/gin"
)

func RegisterWebhooks(rg *gin.RouterGroup, store store.Store) {
	webhooks := rg.Group("webhooks")

	shortcutRG := webhooks.Group("shortcut")
	shortcutRG.POST("/v1", shortcut.WebhooksHandler(NewProcessor("shortcut", store)))

	githubRG := webhooks.Group("github")
	githubRG.POST("/v1", github.WebhooksHandler(NewProcessor("github", store)))
}
