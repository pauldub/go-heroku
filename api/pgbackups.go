package api

import "net/http"

func ListBackups() Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/backups"), nil)
	}
}

func GetBackup(name string) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/backups/"+name), nil)
	}
}

func GetLatestBackup() Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("GET", ApiUrl(apiEndpoint, "/latest_backup"), nil)
	}
}
