package api

import "net/http"
import "go-heroku/types"

func ListAllAddons() Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/addons"), nil)
	}
}

func ListAddons(app *types.Application) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/addons"), nil)
	}
}

func InstallAddon(app *types.Application, addon *types.Addon) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("POST", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/addons/"+addon.Name), nil)
	}
}

func UpgradeAddon(app *types.Application, addon *types.Addon) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("PUT", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/addons/"+addon.Name), nil)
	}
}

func UninstallAddon(app *types.Application, addon *types.Addon) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/apps/"+app.Name+"/addons/"+addon.Name), nil)
	}
}
