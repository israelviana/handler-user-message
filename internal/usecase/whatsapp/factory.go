package whatsapp

import "meta-integration/internal/service"

type whatsappUseCase struct {
	service service.IWhatsappService
}

func NewWhatsappUseCase(svc service.IWhatsappService) IWhatsappUseCase {
	return &whatsappUseCase{
		service: svc,
	}
}
