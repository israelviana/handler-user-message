package clients

import (
	"context"
	"encoding/json"
	"time"

	"handler-user-message/internal/domain/clients/tito"

	"github.com/go-resty/resty/v2"
)

type titoClient struct {
	resty *resty.Client
}

func NewTitoClient(url, apiKey string) tito.ITitoClient {
	return &titoClient{
		resty: resty.New().
			SetBaseURL(url).
			SetHeader("X-CMC_PRO_API_KEY", apiKey).
			SetHeader("Accept", "application/json").
			SetTimeout(15 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(500 * time.Millisecond).
			SetRetryMaxWaitTime(5 * time.Second),
	}
}

func (c *titoClient) SendMessage(ctx context.Context, message string) (interface{}, error) {
	reqBody, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	var response interface{}
	_, err = c.resty.R().SetContext(ctx).SetBody(&reqBody).SetResult(&response).Post("/chat")
	if err != nil {
		return nil, err
	}

	return response, nil
}
