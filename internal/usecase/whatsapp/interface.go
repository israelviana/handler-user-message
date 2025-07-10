package whatsapp

import (
	"context"
	"meta-integration/internal/domain"
)

type IWhatsappUseCase interface {
	SendWhatsappMessage(ctx context.Context, template domain.MetaSendWhatsappMessageBody) (*domain.MetaSendWhatsappMessageWithTemplateResponse, error)
	CreateWhatsappTemplate(ctx context.Context, template domain.MetaCreateTemplateBody) (*domain.MetaCreateTemplateResponse, error)
	ListWhatsappTemplates(ctx context.Context, options ...domain.MetaListTemplatesOption) (map[string]any, error)
	DeleteWhatsappTemplate(ctx context.Context, body string, options ...domain.MetaDeleteTemplateOption) (bool, error)
	EditWhatsappTemplate(ctx context.Context, body domain.MetaEditTemplateBody) (bool, error)
}
