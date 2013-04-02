package api

import "net/http"
import "go-heroku/types"

func ListKeys() Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("POST", ApiUrl(apiEndpoint, "/user/keys"), nil)
	}
}

func AddKey(key *types.SshKey) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		req, err = http.NewRequest("POST", ApiUrl(apiEndpoint, "/user/keys"), nil)
		if err != nil {
			return nil, err
		}

		// TODO: Add ssh key as request body.

		return req, nil
	}
}

func DeleteKey(key *types.SshKey) Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/user/keys/"+key.Email), nil)
	}
}

func DeleteAllKeys() Request {
	return func(apiEndpoint string) (req *http.Request, err error) {
		return http.NewRequest("DELETE", ApiUrl(apiEndpoint, "/user/keys"), nil)
	}
}
