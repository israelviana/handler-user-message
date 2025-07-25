package main

import (
	"meta-integration/cmd/api/rest"
	"meta-integration/cmd/api/webhook"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	wbHandler := webhook.NewWebhookHandler(router)
	wbHandler.InitWebhookHandler()

	rHandler := rest.NewRestHandler(router)
	rHandler.InitRestHandler()
	err := router.Run(":50051")
	if err != nil {
		return
	}
	/*
		baseURL := os.Getenv("BASE_URL_META")
			accessToken := os.Getenv("ACCESS_TOKEN_META")
			fromNumberID := os.Getenv("FROM_NUMBER_ID_META")
			whatsappBusinessAccountID := os.Getenv("WHATSAPP_BUSINESS_ACCOUNT_ID_META")

			client := resty.New()

		newWpService := service.NewWhatsappService(client, baseURL, accessToken, fromNumberID, whatsappBusinessAccountID)
		newWpUseCase := wpUseCase.NewWhatsappUseCase(newWpService)


		handler := grpcHandler.NewGrpcHandler(newWpUseCase)

		server := grpc.NewServer()
		pb.RegisterWhatsappServiceServer(server, handler)

		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("❌ Failed to listen on port 50051: %v", err)
		}
		reflection.Register(server)
		log.Println("✅ gRPC server running on :50051")
		if err := server.Serve(listener); err != nil {
			log.Fatalf("❌ Failed to serve gRPC: %v", err)
		}*/

}
