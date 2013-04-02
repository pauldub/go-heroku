package api

import _ "fmt"
import "net/http"

type Request func(apiEndpoint string) (*http.Request, error)

func ApiUrl(baseUrl string, methodPath string) string {
	return baseUrl + methodPath
}
