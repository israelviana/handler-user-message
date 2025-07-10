package whatsapp

import (
	"context"
	"errors"
	"meta-integration/internal/domain"
)

func (uc *whatsappUseCase) CreateWhatsappTemplate(ctx context.Context, body domain.MetaCreateTemplateBody) (*domain.MetaCreateTemplateResponse, error) {
	if body.Name == "" || body.Category == "" {
		return nil, errors.New("missing required fields")
	}
	return uc.service.CreateWhatsappTemplate(ctx, body)
}
