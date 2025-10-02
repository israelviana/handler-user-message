package webhook

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"handler-user-message/internal/domain/clients/whatsapp"
	"handler-user-message/internal/usecase"

	"github.com/gin-gonic/gin"
)

type IWebhookHandler interface {
	InitWebhookHandler()
	VerifyWebhook(c *gin.Context)
	HandleWebhook(c *gin.Context)
}

type WebhookHandler struct {
	router                 *gin.Engine
	processIncomingMessage *usecase.ProcessIncomingMessageUseCase
}

func NewWebhookHandler(router *gin.Engine, processIncomingMessageUC *usecase.ProcessIncomingMessageUseCase) IWebhookHandler {
	return &WebhookHandler{
		router:                 router,
		processIncomingMessage: processIncomingMessageUC,
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
		fmt.Println("WEBHOOK_VERIFIED")
		c.String(http.StatusOK, challenge)
		return
	}

	c.AbortWithStatus(http.StatusForbidden)
}

func (w *WebhookHandler) HandleWebhook(c *gin.Context) {
	var payload whatsapp.MetaWebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	for _, entry := range payload.Entry {
		for _, change := range entry.Changes {
			switch change.Field {
			case "messages":
				ch := change
				ctx := c.Copy()

				go func(ch whatsapp.Change, ctx *gin.Context) {
					defer func() {
						if r := recover(); r != nil {
							fmt.Println("panic recovered in processIncomingMessage:", r)
						}
					}()

					jobCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()

					for index, message := range ch.Value.Messages {
						if err := w.processIncomingMessage.Run(jobCtx, message.Text.Body, change.Value.Contacts[index].WAID); err != nil {
							log.Println("error processing incoming message:", err)
						}
					}

				}(ch, ctx)
				continue
			case "statuses":
				log.Println(change.Value)
			default:
				log.Println(change.Value)
			}
		}
	}

	c.Status(http.StatusOK)
}
