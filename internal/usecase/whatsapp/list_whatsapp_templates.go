package whatsapp

import (
	"context"
	"meta-integration/internal/domain"
)

func (uc *whatsappUseCase) ListWhatsappTemplates(ctx context.Context, options ...domain.MetaListTemplatesOption) (map[string]any, error) {
	return uc.service.ListWhatsappTemplates(ctx, options...)
}
