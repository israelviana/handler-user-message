package usecase

import (
	"context"
	"fmt"
	"log"

	"handler-user-message/internal/domain/clients/tito"
	"handler-user-message/internal/domain/clients/whatsapp"
	_ "handler-user-message/internal/domain/clients/whatsapp"
)

type iTitoClient interface {
	SendMessage(ctx context.Context, message string) (interface{}, error)
}

type iWhatsappClient interface {
	SendWhatsappMessage(ctx context.Context, template whatsapp.MetaSendWhatsappMessageBody) (*whatsapp.MetaSendWhatsappMessageWithTemplateResponse, error)
}

type ProcessIncomingMessageUseCase struct {
	titoClient     iTitoClient
	whatsappClient iWhatsappClient
}

func NewProcessIncomingMessageUseCase(titoClient tito.ITitoClient, whatsappClient whatsapp.IWhatsappClient) *ProcessIncomingMessageUseCase {
	return &ProcessIncomingMessageUseCase{
		titoClient:     titoClient,
		whatsappClient: whatsappClient,
	}
}

func (uc *ProcessIncomingMessageUseCase) Run(ctx context.Context, message string) error {
	res, err := uc.titoClient.SendMessage(ctx, message)
	if err != nil {
		log.Println(err)
		return err
	}

	s := fmt.Sprintf("%v", res)
	_, err = uc.whatsappClient.SendWhatsappMessage(ctx, whatsapp.MetaSendWhatsappMessageBody{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               "85997267265",
		Type:             "text",
		Text: &whatsapp.MessageText{
			PreviewUrl: false,
			Body:       s,
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
