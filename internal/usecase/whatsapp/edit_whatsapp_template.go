package whatsapp

import (
	"context"
	"meta-integration/internal/domain"
)

func (uc *whatsappUseCase) EditWhatsappTemplate(ctx context.Context, body domain.MetaEditTemplateBody) (bool, error) {
	return uc.service.EditWhatsappTemplate(ctx, body)
}
