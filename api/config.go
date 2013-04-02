package api

import "net/http"
import "go-heroku/types"

func ListConfig(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/config_vars"), nil)
	}
}

func AddConfig(app *types.Application, configVars *[]types.Config) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("PUT", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/config_vars"), nil)
		if err != nil {
			return nil, err
		}

		//TODO: Marshal configVars into JSON as request body.

		return req, nil
	}
}

func DeleteConfig(app *types.Application, config *types.Config) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/config_vars/"+config.Key), nil)
	}
}
