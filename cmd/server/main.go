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
	baseURL := getEnvOrFail("BASE_URL_META")
	accessToken := getEnvOrFail("ACCESS_TOKEN_META")
	fromNumberID := getEnvOrFail("FROM_NUMBER_ID_META")
	whatsappBusinessAccountID := getEnvOrFail("WHATSAPP_BUSINESS_ACCOUNT_ID_META")

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

func getEnvOrFail(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("❌ Required environment variable %s is missing", key)
	}
	return val
}
