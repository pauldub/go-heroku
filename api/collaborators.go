package api

import "net/http"
import "net/url"
import "go-heroku/types"

func ListCollaborators(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/collaborators"), nil)
	}
}

func AddCollaborator(app *types.Application, collaborator *types.Collaborator) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/collaborators"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("collaborator[email]", collaborator.Email)

		req.Form = params

		return req, nil
	}
}

func RemoveCollaborator(app *types.Application, collaborator *types.Collaborator) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/collaborators/"+collaborator.Email), nil)
	}
}
