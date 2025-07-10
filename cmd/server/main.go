package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
	wpUseCase "meta-integration/internal/usecase/whatsapp"
	"net"
	"os"

	"google.golang.org/grpc"
	cmdGrpc "meta-integration/cmd/api/grpc"
	pb "meta-integration/gen/proto"
	"meta-integration/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL, isValid := os.LookupEnv("BASE_URL_META")
	if !isValid {
		log.Fatal("Error loading BASE_URL_META env")
	}
	_, isValid = os.LookupEnv("ACCESS_TOKEN_META")
	if !isValid {
		log.Fatal("Error loading ACCESS_TOKEN_META env")
	}
	accessToken, isValid := os.LookupEnv("BACKUP_ACCESS_TOKEN_META")
	if !isValid {
		log.Fatal("Error loading BACKUP_ACCESS_TOKEN_META env")
	}
	fromNumberID, isValid := os.LookupEnv("FROM_NUMBER_ID_META")
	if !isValid {
		log.Fatal("Error loading FROM_NUMBER_META env")
	}
	whatsappBusinessAccountId, isValid := os.LookupEnv("WHATSAPP_BUSINESS_ACCOUNT_ID_META")
	if !isValid {
		log.Fatal("Error loading WHATSAPP_BUSINESS_ACCOUNT_ID_META env")
	}

	restyClient := resty.New()
	newWpService := service.NewWhatsappService(restyClient, baseURL, accessToken, fromNumberID, whatsappBusinessAccountId)
	newWpUseCase := wpUseCase.NewWhatsappUseCase(newWpService)
	newGrpcHandler := cmdGrpc.NewGrpcHandler(newWpUseCase)

	grpcServer := grpc.NewServer()
	pb.RegisterWhatsappServiceServer(grpcServer, newGrpcHandler)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gRPC server running on :50051")
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
