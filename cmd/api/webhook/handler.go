package webhook

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meta-integration/internal/domain"
	"net/http"
	"os"
)

type IWebhookHandler interface {
	InitWebhookHandler()
	VerifyWebhook(c *gin.Context)
	HandleWebhook(c *gin.Context)
}

type WebhookHandler struct {
	router *gin.Engine
}

func NewWebhookHandler(router *gin.Engine) IWebhookHandler {
	return &WebhookHandler{
		router: router,
	}
}

func (w *WebhookHandler) InitWebhookHandler() {
	w.router.GET("/webhook", w.VerifyWebhook)
	w.router.POST("/webhook", w.HandleWebhook)
}

func (w *WebhookHandler) VerifyWebhook(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	expectedToken := os.Getenv("META_VERIFY_TOKEN")

	if mode == "subscribe" && token == expectedToken {
		c.String(http.StatusOK, challenge)
		return
	}

	c.AbortWithStatus(http.StatusForbidden)
}

func (w *WebhookHandler) HandleWebhook(c *gin.Context) {
	var payload domain.MetaWebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	for _, entry := range payload.Entry {
		for _, change := range entry.Changes {
			switch change.Field {
			case "messages":
				fmt.Println(change.Value)
			case "statuses":
				fmt.Println(change.Value)
			default:

			}
		}
	}

	c.Status(http.StatusOK)
}
