package main

import (
	"os"

	"meta-integration/clients"
	"meta-integration/cmd/api/rest"
	"meta-integration/cmd/api/webhook"
	"meta-integration/internal/usecase/tito"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	baseURL := os.Getenv("BASE_URL_META")
	accessToken := os.Getenv("ACCESS_TOKEN_META")
	fromNumberID := os.Getenv("FROM_NUMBER_ID_META")
	whatsappBusinessAccountID := os.Getenv("WHATSAPP_BUSINESS_ACCOUNT_ID_META")

	router := gin.Default()

	titoClient := clients.NewTitoClient(baseURL, accessToken)
	whatsappClient := clients.NewWhatsappClient(resty.New(), baseURL, accessToken, fromNumberID, whatsappBusinessAccountID)
	processIncomingMessageUseCase := tito.NewProcessIncomingMessageUseCase(titoClient, whatsappClient)
	wHandler := webhook.NewWebhookHandler(router, processIncomingMessageUseCase)
	wHandler.InitWebhookHandler()

	rHandler := rest.NewRestHandler(router)
	rHandler.InitRestHandler()

	err := router.Run(":50051")
	if err != nil {
		return
	}
}
