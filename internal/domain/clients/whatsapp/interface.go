package whatsapp

import (
	"context"
)

type IWhatsappClient interface {
	SendWhatsappMessage(ctx context.Context, template MetaSendWhatsappMessageBody) (*MetaSendWhatsappMessageWithTemplateResponse, error)
	CreateWhatsappTemplate(ctx context.Context, template MetaCreateTemplateBody) (*MetaCreateTemplateResponse, error)
	ListWhatsappTemplates(ctx context.Context, options ...MetaListTemplatesOption) (map[string]any, error)
	DeleteWhatsappTemplate(ctx context.Context, body string, options ...MetaDeleteTemplateOption) (bool, error)
	EditWhatsappTemplate(ctx context.Context, body MetaEditTemplateBody) (bool, error)
}
