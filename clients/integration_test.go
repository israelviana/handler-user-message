package clients

/*
import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"meta-integration/internal/domain/clients/whatsapp"
	"meta-integration/internal/service"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func TestWhatsappService(t *testing.T) {
	ctx := context.Background()

	err := godotenv.Load("../../.env")
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
	wpR := service.NewWhatsappService(restyClient, baseURL, accessToken, fromNumberID, whatsappBusinessAccountId)

	message := whatsapp.MetaSendWhatsappMessageBody{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               "+5585997267265",
		Type:             "template",
		Template: &whatsapp.MetaTemplate{
			Name:     "hello_world",
			Language: whatsapp.MetaLanguage{Code: "en_US"},
		},
	}
	response, err := wpR.SendWhatsappMessage(ctx, message)
	if err != nil {
		log.Fatal(fmt.Sprintf("%v, details: %s", err, err.Error()))
	}

	log.Println(response)
}
*/
