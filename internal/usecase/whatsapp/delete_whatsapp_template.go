package whatsapp

import (
	"context"
	"meta-integration/internal/domain"
)

func (uc *whatsappUseCase) DeleteWhatsappTemplate(ctx context.Context, body string, options ...domain.MetaDeleteTemplateOption) (bool, error) {
	return uc.service.DeleteWhatsappTemplate(ctx, body, options...)
}
