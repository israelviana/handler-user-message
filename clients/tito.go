package clients

import (
	"context"
	"encoding/json"
	"fmt"
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

	req, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(req))
	fmt.Println(c.resty.BaseURL)

	var response interface{}
	_, err = c.resty.R().SetContext(ctx).SetBody(&req).SetResult(&response).Post("/chat")
	if err != nil {
		return nil, err
	}

	return response, nil
}
