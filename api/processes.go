package api

import "fmt"
import "net/http"
import "net/url"
import "go-heroku/types"

func ListProcesses(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/ps"), nil)
	}
}

func RunProcess(app *types.Application, command string) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/ps"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("command", command)

		req.Form = params

		return req, nil
	}
}

func RestartProcess(app *types.Application, processName string, processType string) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/ps/restart"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("ps", processName)
		params.Add("type", processType)

		req.Form = params

		return req, nil
	}
}

func StopProcess(app *types.Application, processName string, processType string) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/ps/stop"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("ps", processName)
		params.Add("type", processType)

		req.Form = params

		return req, nil
	}
}

func ScaleProcess(app *types.Application, processType string, quantity int) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/ps/scale"), nil)
		if err != nil {
			return nil, err
		}

		params := url.Values{}
		params.Add("type", processType)
		params.Add("qte", fmt.Sprintf("%v", quantity))

		req.Form = params

		return req, nil
	}
}
