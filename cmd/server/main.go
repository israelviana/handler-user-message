package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
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
	if os.Getenv("RUNNING_IN_DOCKER") != "true" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ Alert: .env dont have loaded.")
		}
	}

	baseURL := getEnvOrFail("BASE_URL_META")
	accessToken := getEnvOrFail("ACCESS_TOKEN_META")
	fromNumberID := getEnvOrFail("FROM_NUMBER_ID_META")
	whatsappBusinessAccountID := getEnvOrFail("WHATSAPP_BUSINESS_ACCOUNT_ID_META")

	client := resty.New()
	newWpService := service.NewWhatsappService(client, baseURL, accessToken, fromNumberID, whatsappBusinessAccountID)
	newWpUseCase := wpUseCase.NewWhatsappUseCase(newWpService)
	newGrpcHandler := grpcHandler.NewGrpcHandler(newWpUseCase)

	grpcServer := grpc.NewServer()
	pb.RegisterWhatsappServiceServer(grpcServer, newGrpcHandler)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func getEnvOrFail(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("❌ Env not found: %s", key)
	}
	return val
}
