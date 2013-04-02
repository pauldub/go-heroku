package client

import (
	"encoding/base64"
	"fmt"
	"go-heroku/api"
	"net/http"
)

type HerokuClient struct {
	ApiKey      string
	ApiEndpoint string

	httpClient *http.Client
}

func NewHerokuClient(apiKey string) *HerokuClient {
	httpClient := &http.Client{}
	return &HerokuClient{apiKey, "https://api.heroku.com", httpClient}
}

func NewHerokuClientFromEnv() *HerokuClient {
	// TODO: Load apiKey from env
	httpClient := &http.Client{}
	return &HerokuClient{"", "https://api.heroku.com", httpClient}
}

func NewHerokuClientFromFile(path string) *HerokuClient {
	// TODO: Load apiKey from file
	httpClient := &http.Client{}
	return &HerokuClient{"", "https://api.heroku.com", httpClient}
}

func (client *HerokuClient) authorizationHeader() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Basic :%v", client.ApiKey)))
}

func (client *HerokuClient) setHeaders(headers *http.Header) *http.Header {
	headers.Add("Authorization", client.authorizationHeader())
	headers.Add("Accept", "application/json")
	return headers
}

func (client *HerokuClient) Do(apiRequest api.Request, data interface{}) error {
	req, err := apiRequest(client.ApiEndpoint)
	if err != nil {
		return err
	}

	client.setHeaders(&req.Header)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}

	err = decodeApiResponse(res, data)
	if err != nil {
		return err
	}

	return nil
}
