package whatsapp

import (
	"context"
	"meta-integration/internal/domain"
)

func (uc *whatsappUseCase) SendWhatsappMessage(ctx context.Context, template domain.MetaSendWhatsappMessageBody) (*domain.MetaSendWhatsappMessageWithTemplateResponse, error) {
	return uc.service.SendWhatsappMessage(ctx, template)
}
