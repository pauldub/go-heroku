package api

import "net/http"
import "net/url"
import "go-heroku/types"

func ListApplications(apiEndpoint string) (req *http.Request, err error) {
	return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps"), nil)
}

func CreateApplication(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("app[name]", app.Name)

		req.Form = params

		return req, nil
	}
}

func RenameApplication(app *types.Application, newName string) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("PUT", ApiUrl(apiEndpoint, "/apps/"+app.Name), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("app[name]", newName)

		req.Form = params

		return req, nil
	}
}

func TransferApplication(app *types.Application, collaborator *types.Collaborator) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("PUT", ApiUrl(apiEndpoint, "/apps/"+app.Name), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("app[transfer_owner]", collaborator.Email)

		req.Form = params

		return req, nil
	}
}

func SetApplicationMaintenance(app *types.Application, active bool) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		if active == true {
			params.Add("maintenance_mode", "1")
		} else {
			params.Add("maintenance_mode", "0")
		}

		req.Form = params

		return req, nil
	}
}

func DeleteApplication(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/apps/"+app.Name), nil)
	}
}
