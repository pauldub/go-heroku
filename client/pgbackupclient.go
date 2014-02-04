package client

import "fmt"
import "errors"
import "net/http"
import "go-heroku/api"
import "go-heroku/types"

// import "io/ioutil"

type PgBackupClient struct {
	App         *types.Application
	Addon       *types.Addon
	httpClient  *http.Client
	ApiEndpoint string
}

/* PgBackups */

const PgBackupUrlKey = "PGBACKUPS_URL"

func NewPgBackupClient(app *types.Application, addon *types.Addon, apiUrl string) *PgBackupClient {
	httpClient := &http.Client{}
	return &PgBackupClient{app, addon, httpClient, apiUrl}
}

func NewPgBackupClientFromHerokuClient(herokuClient *HerokuClient, app *types.Application, addon *types.Addon) (*PgBackupClient, error) {
	configs, err := herokuClient.ListConfig(app)
	if err != nil {
		return nil, err
	}
	config := findPgBackupUrl(configs)
	if config == nil {
		message := fmt.Sprintf("could not find %s in config", PgBackupUrlKey)
		return nil, errors.New(message)
	}

	return NewPgBackupClient(app, addon, config.Value), nil
}

func (client *PgBackupClient) Do(apiRequest api.Request, data interface{}) error {
	req, err := apiRequest(client.ApiEndpoint)
	if err != nil {
		return err
	}

	client.setHeaders(&req.Header)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		return &ApiError{Message: "Object not found"}
	}
	err = decodeApiResponse(res, data)
	if err != nil {
		return err
	}

	return nil
}

/* private */

func (client *PgBackupClient) setHeaders(headers *http.Header) *http.Header {
	headers.Add("X-Heroku-Gem-Version", "3")
	return headers
}

func findPgBackupUrl(configs *[]types.Config) *types.Config {
	for _, config := range *configs {
		if config.Key == PgBackupUrlKey {
			return &config
		}
	}
	return nil
}
