package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"handler-user-message/internal/domain/clients/whatsapp"

	"github.com/go-resty/resty/v2"
)

type whatsappClient struct {
	client                    *resty.Client
	baseURL                   string
	accessToken               string
	fromNumberID              string
	whatsappBusinessAccountID string
}

func NewWhatsappClient(client *resty.Client, baseURL, accessToken, fromNumberID, whatsappBusinessAccountId string) whatsapp.IWhatsappClient {
	return &whatsappClient{
		client:                    client,
		baseURL:                   baseURL,
		accessToken:               accessToken,
		fromNumberID:              fromNumberID,
		whatsappBusinessAccountID: whatsappBusinessAccountId,
	}
}

func (r *whatsappClient) SendWhatsappMessage(ctx context.Context, message whatsapp.MetaSendWhatsappMessageBody) (*whatsapp.MetaSendWhatsappMessageWithTemplateResponse, error) {
	///<WHATSAPP_BUSINESS_PHONE_NUMBER_ID>/messages
	url := fmt.Sprintf("%s/%s/messages", r.baseURL, r.fromNumberID)

	resp := new(whatsapp.MetaSendWhatsappMessageWithTemplateResponse)
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

func (r *whatsappClient) CreateWhatsappTemplate(ctx context.Context, body whatsapp.MetaCreateTemplateBody) (*whatsapp.MetaCreateTemplateResponse, error) {
	// /<WHATSAPP_BUSINESS_ACCOUNT_ID>/message_templates
	url := fmt.Sprintf("%s/%s/message_templates", r.baseURL, r.whatsappBusinessAccountID)

	resp := new(whatsapp.MetaCreateTemplateResponse)
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

func (r *whatsappClient) ListWhatsappTemplates(ctx context.Context, options ...whatsapp.MetaListTemplatesOption) (map[string]any, error) {
	// /<WHATSAPP_BUSINESS_ACCOUNT_ID>/message_templates
	url := fmt.Sprintf("%s/%s/message_templates", r.baseURL, r.whatsappBusinessAccountID)

	params := &whatsapp.MetaListTemplatesParams{}
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

func (r *whatsappClient) DeleteWhatsappTemplate(ctx context.Context, templateName string, options ...whatsapp.MetaDeleteTemplateOption) (bool, error) {
	url := fmt.Sprintf("%s/%s/message_templates", r.baseURL, r.whatsappBusinessAccountID)

	params := &whatsapp.MetaDeleteTemplateParams{}
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

func (r *whatsappClient) EditWhatsappTemplate(ctx context.Context, body whatsapp.MetaEditTemplateBody) (bool, error) {
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
