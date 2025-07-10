package grpc

import (
	"context"
	metaPb "meta-integration/gen/proto"
	"meta-integration/internal/domain"
	wpUsecase "meta-integration/internal/usecase/whatsapp"
	"meta-integration/utils"
)

type GrpcHandler struct {
	metaPb.UnimplementedWhatsappServiceServer
	whatsappUseCase wpUsecase.IWhatsappUseCase
}

func NewGrpcHandler(whatsappUseCase wpUsecase.IWhatsappUseCase) *GrpcHandler {
	return &GrpcHandler{whatsappUseCase: whatsappUseCase}
}

func (h *GrpcHandler) SendWhatsappMessage(ctx context.Context, req *metaPb.MetaSendWhatsappMessageBody) (*metaPb.MetaSendWhatsappMessageWithTemplateResponse, error) {
	dto, err := utils.JsonMarshalToUnmarshal[domain.MetaSendWhatsappMessageBody](req)
	if err != nil {
		return nil, err
	}

	response, err := h.whatsappUseCase.SendWhatsappMessage(ctx, *dto)
	if err != nil {
		return nil, err
	}

	proto, err := utils.JsonMarshalToUnmarshal[metaPb.MetaSendWhatsappMessageWithTemplateResponse](response)
	if err != nil {
		return nil, err
	}

	return proto, nil
}

func (h *GrpcHandler) CreateWhatsappTemplate(ctx context.Context, req *metaPb.MetaCreateTemplateBody) (*metaPb.MetaCreateTemplateResponse, error) {
	dto, err := utils.JsonMarshalToUnmarshal[domain.MetaCreateTemplateBody](req)
	if err != nil {
		return nil, err
	}

	response, err := h.whatsappUseCase.CreateWhatsappTemplate(ctx, *dto)
	if err != nil {
		return nil, err
	}

	proto, err := utils.JsonMarshalToUnmarshal[metaPb.MetaCreateTemplateResponse](response)
	if err != nil {
		return nil, err
	}

	return proto, nil
}

func (h *GrpcHandler) ListWhatsappTemplates(ctx context.Context, req *metaPb.MetaListTemplatesParams) (*metaPb.MetaListTemplatesResponse, error) {
	dto, err := utils.JsonMarshalToUnmarshal[domain.MetaListTemplatesParams](req)
	if err != nil {
		return nil, err
	}

	var options []domain.MetaListTemplatesOption
	if dto.Limit != 0 {
		domain.WithLimit(dto.Limit)
	}

	if len(dto.Fields) > 0 {
		domain.WithFields(dto.Fields)
	}

	response, err := h.whatsappUseCase.ListWhatsappTemplates(ctx, options...)
	if err != nil {
		return nil, err
	}

	proto, err := utils.JsonMarshalToUnmarshal[metaPb.MetaListTemplatesResponse](response)
	if err != nil {
		return nil, err
	}

	return proto, nil
}

func (h *GrpcHandler) DeleteWhatsappTemplate(ctx context.Context, req *metaPb.MetaDeleteTemplateParams) (*metaPb.MetaDeleteWhatsappTemplateResponse, error) {
	dto, err := utils.JsonMarshalToUnmarshal[domain.MetaDeleteTemplateParams](req)
	if err != nil {
		return nil, err
	}

	var options []domain.MetaDeleteTemplateOption
	if dto.ID != "" {
		domain.WithID(dto.ID)
	}

	response, err := h.whatsappUseCase.DeleteWhatsappTemplate(ctx, req.GetTemplateName(), options...)
	if err != nil {
		return nil, err
	}

	proto, err := utils.JsonMarshalToUnmarshal[metaPb.MetaDeleteWhatsappTemplateResponse](response)
	if err != nil {
		return nil, err
	}

	return proto, nil
}

func (h *GrpcHandler) EditWhatsappTemplate(ctx context.Context, req *metaPb.MetaEditTemplateBody) (*metaPb.MetaEditTemplateResponse, error) {
	dto, err := utils.JsonMarshalToUnmarshal[domain.MetaEditTemplateBody](req)
	if err != nil {
		return nil, err
	}

	response, err := h.whatsappUseCase.EditWhatsappTemplate(ctx, *dto)
	if err != nil {
		return nil, err
	}

	proto, err := utils.JsonMarshalToUnmarshal[metaPb.MetaEditTemplateResponse](response)
	if err != nil {
		return nil, err
	}

	return proto, nil
}
