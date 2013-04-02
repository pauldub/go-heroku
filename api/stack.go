package api

import "net/http"
import "go-heroku/types"

func ListStacks(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/stack"), nil)
	}
}

func MigrateStack(app *types.Application, stack *types.Stack) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("PUT", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/stack"), nil)
		if err != nil {
			return nil, err
		}

		// TODO: Add stack as request body.

		return req, nil
	}
}
