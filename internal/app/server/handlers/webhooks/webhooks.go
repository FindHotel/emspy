package webhooks

import (
	"github.com/FindHotel/emspy/internal/app/server/handlers/webhooks/shortcut"
	"github.com/FindHotel/emspy/internal/app/server/store"
	"github.com/gin-gonic/gin"
)

func RegisterWebhooks(rg *gin.RouterGroup, store store.Store) {
	webhooks := rg.Group("webhooks")

	shortcutRG := webhooks.Group("shortcut")

	shortcutRG.POST("/v1", shortcut.WebhooksHandler(NewProcessor(store)))
}
