package main

import (
	wpUsecase "meta-integration/internal/usecase/whatsapp"
)

type RestHandler struct {
	whatsappUseCase wpUsecase.IWhatsappUseCase
}

func NewRestHandler(whatsappUseCase wpUsecase.IWhatsappUseCase) *RestHandler {
	return &RestHandler{whatsappUseCase: whatsappUseCase}
}
