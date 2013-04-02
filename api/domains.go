package api

import "net/http"
import "net/url"
import "go-heroku/types"

func ListDomains(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/domains"), nil)
	}
}

func AddDomain(app *types.Application, domain *types.Domain) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/domains"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("domain_name[domain]", domain.Domain)

		req.Form = params

		return req, nil
	}
}

func DeleteDomain(app *types.Application, domain *types.Domain) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/domains/"+domain.Domain), nil)
	}
}
