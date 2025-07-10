package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"meta-integration/internal/domain"
	"strings"
)

type whatsappService struct {
	client                    *resty.Client
	baseURL                   string
	accessToken               string
	fromNumberID              string
	whatsappBusinessAccountID string
}

type IWhatsappService interface {
	SendWhatsappMessage(ctx context.Context, template domain.MetaSendWhatsappMessageBody) (*domain.MetaSendWhatsappMessageWithTemplateResponse, error)
	CreateWhatsappTemplate(ctx context.Context, template domain.MetaCreateTemplateBody) (*domain.MetaCreateTemplateResponse, error)
	ListWhatsappTemplates(ctx context.Context, options ...domain.MetaListTemplatesOption) (map[string]any, error)
	DeleteWhatsappTemplate(ctx context.Context, body string, options ...domain.MetaDeleteTemplateOption) (bool, error)
	EditWhatsappTemplate(ctx context.Context, body domain.MetaEditTemplateBody) (bool, error)
}

func NewWhatsappService(client *resty.Client, baseURL, accessToken, fromNumberID, whatsappBusinessAccountId string) IWhatsappService {
	return &whatsappService{
		client:                    client,
		baseURL:                   baseURL,
		accessToken:               accessToken,
		fromNumberID:              fromNumberID,
		whatsappBusinessAccountID: whatsappBusinessAccountId,
	}
}

func (r *whatsappService) SendWhatsappMessage(ctx context.Context, message domain.MetaSendWhatsappMessageBody) (*domain.MetaSendWhatsappMessageWithTemplateResponse, error) {
	// /<WHATSAPP_BUSINESS_PHONE_NUMBER_ID>/messages
	url := fmt.Sprintf("%s/%s/messages", r.baseURL, r.fromNumberID)

	resp := new(domain.MetaSendWhatsappMessageWithTemplateResponse)
	httpResp, err := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+r.accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(message).
		SetResult(resp).
		Post(url)

	if err != nil {
		return nil, err
	}
	if httpResp.IsError() {
		return nil, fmt.Errorf("meta error: %s", httpResp.Status())
	}

	return resp, nil
}

func (r *whatsappService) CreateWhatsappTemplate(ctx context.Context, body domain.MetaCreateTemplateBody) (*domain.MetaCreateTemplateResponse, error) {
	// /<WHATSAPP_BUSINESS_ACCOUNT_ID>/message_templates
	url := fmt.Sprintf("%s/%s/message_templates", r.baseURL, r.whatsappBusinessAccountID)

	resp := new(domain.MetaCreateTemplateResponse)
	httpResp, err := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+r.accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(resp).
		Post(url)

	if err != nil {
		return nil, err
	}
	if httpResp.IsError() {
		return nil, fmt.Errorf("meta error: %s", httpResp.Status())
	}

	return resp, nil
}

func (r *whatsappService) ListWhatsappTemplates(ctx context.Context, options ...domain.MetaListTemplatesOption) (map[string]any, error) {
	// /<WHATSAPP_BUSINESS_ACCOUNT_ID>/message_templates
	url := fmt.Sprintf("%s/%s/message_templates", r.baseURL, r.whatsappBusinessAccountID)

	params := &domain.MetaListTemplatesParams{}
	for _, opt := range options {
		opt(params)
	}

	req := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+r.accessToken).
		SetHeader("Content-Type", "application/json")

	if len(params.Fields) > 0 {
		req.SetQueryParam("fields", strings.Join(params.Fields, ","))
	}
	if params.Limit != 0 {
		req.SetQueryParam("limit", fmt.Sprintf("%d", params.Limit))
	}

	httpResp, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	if httpResp.IsError() {
		return nil, fmt.Errorf("meta error: %s", httpResp.Status())
	}

	var resp map[string]any
	err = json.Unmarshal(httpResp.Body(), &resp)
	if err != nil {
		return nil, fmt.Errorf("error to unmarhall response: %w", err)
	}

	return resp, nil
}

func (r *whatsappService) DeleteWhatsappTemplate(ctx context.Context, templateName string, options ...domain.MetaDeleteTemplateOption) (bool, error) {
	url := fmt.Sprintf("%s/%s/message_templates", r.baseURL, r.whatsappBusinessAccountID)

	params := &domain.MetaDeleteTemplateParams{}
	for _, opt := range options {
		opt(params)
	}

	req := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+r.accessToken).
		SetHeader("Content-Type", "application/json").
		SetQueryParam("name", templateName)

	if params.ID != "" {
		req.SetQueryParam("hsm_id", params.ID)
	}

	httpResp, err := req.Delete(url)
	if err != nil {
		return false, err
	}
	if httpResp.IsError() {
		return false, fmt.Errorf("meta error: %s", httpResp.Status())
	}

	return httpResp.IsSuccess(), nil
}

func (r *whatsappService) EditWhatsappTemplate(ctx context.Context, body domain.MetaEditTemplateBody) (bool, error) {
	// /<WHATSAPP_MESSAGE_TEMPLATE_ID>

	url := fmt.Sprintf("%s/%s", r.baseURL, r.whatsappBusinessAccountID)

	httpResp, err := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+r.accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if err != nil {
		return false, err
	}
	if httpResp.IsError() {
		return false, fmt.Errorf("meta error: %s", httpResp.Status())
	}

	return httpResp.IsSuccess(), nil
}
