package tito

import (
	"context"

	"meta-integration/internal/domain/clients/tito"
	"meta-integration/internal/domain/clients/whatsapp"
	_ "meta-integration/internal/domain/clients/whatsapp"
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
	_, err := uc.titoClient.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	_, err = uc.whatsappClient.SendWhatsappMessage(ctx, whatsapp.MetaSendWhatsappMessageBody{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               "85997267265",
		Type:             "text",
		Text: &whatsapp.MessageText{
			PreviewUrl: false,
			Body:       "Respondendo com sucesso",
		},
	})
	if err != nil {
		return err
	}

	return nil
}
