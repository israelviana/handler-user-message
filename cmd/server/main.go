package main

import (
	"github.com/go-resty/resty/v2"
	"log"
	grpcHandler "meta-integration/cmd/api/grpc"
	pb "meta-integration/gen/proto"
	"meta-integration/internal/service"
	wpUseCase "meta-integration/internal/usecase/whatsapp"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
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

	log.Println("✅ gRPC server running on :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}
}
