package api

import "net/http"
import "net/url"
import "go-heroku/types"

func ListReleases(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/releases"), nil)
	}
}

func GetRelease(app *types.Application, release *types.Release) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/releases/"+release.Name), nil)
	}
}

func RollbackToRelease(app *types.Application, release *types.Release) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/releases"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("rollback", release.Name)

		req.Form = params

		return req, nil
	}
}
