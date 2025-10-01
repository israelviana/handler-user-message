package clients

import (
	"context"
	"log"
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
			SetHeader("Accept", "application/json").
			SetTimeout(15 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(500 * time.Millisecond).
			SetRetryMaxWaitTime(5 * time.Second),
	}
}

func (c *titoClient) SendMessage(ctx context.Context, message string) (interface{}, error) {
	type RequestBody struct {
		Input string `json:"input"`
	}

	reqBody := RequestBody{
		Input: message,
	}

	var response interface{}
	_, err := c.resty.R().SetContext(ctx).SetBody(&reqBody).SetResult(&response).Post("/chat")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}
