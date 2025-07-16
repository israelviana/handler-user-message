package webhook

type WebhookUseCase struct {
}

func NewWebhookUseCase() IWebhookUseCase {
	return &WebhookUseCase{}
}
