package usecase

import (
	"context"
	"errors"
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

func (uc *ProcessIncomingMessageUseCase) Run(ctx context.Context, message string, sender string) error {
	res, err := uc.titoClient.SendMessage(ctx, message)
	if err != nil {
		log.Println(err)
		return err
	}

	if res == nil {
		return errors.New("error to call chat bot service")
	}

	m, ok := res.(map[string]interface{})
	if !ok {
		log.Println("chat bot response is not a map[string]interface{]", m)
		return errors.New("res is not a map[string]interface{}")
	}

	answer, ok := m["answer"].(string)
	if !ok {
		log.Println("answer is not a valid string", m)
		return errors.New("answer is not string")
	}

	_, err = uc.whatsappClient.SendWhatsappMessage(ctx, whatsapp.MetaSendWhatsappMessageBody{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               sender,
		Type:             "text",
		Text: &whatsapp.MessageText{
			PreviewUrl: false,
			Body:       answer,
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
