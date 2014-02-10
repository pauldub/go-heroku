package client

import "go-heroku/types"
import "testing"
import "net/http"
import "net/http/httptest"

const JSONConfig = `{
"NEW_RELIC_APP_NAME":"motocab",
"RUBY_VERSION":"ruby-1.9.3-p0",
"SENDGRID_PASSWORD":"xxxx",
"PGBACKUPS_URL":"https://xxx:xxxx@pgbackups.heroku.com/client",
"NEW_RELIC_LICENSE_KEY":"xxxxx",
"DATABASE_URL":"postgres://xxxx:xxxx@ec2-23-21-105-133.compute-1.amazonaws.com:5432/xxxx",
"NEW_RELIC_LOG":"stdout",
"HEROKU_POSTGRESQL_MAROON_URL":"postgres://xxxx:xxxx@ec2-23-21-105-133.compute-1.amazonaws.com:5432/xxxx",
"RAILS_ENV":"staging",
"PATH":"bin:vendor/bundle/ruby/1.9.1/bin:/usr/local/bin:/usr/bin:/bin",
"LANG":"en_US.UTF-8",
"RACK_ENV":"production",
"SENDGRID_USERNAME":"sxxxx@heroku.com",
"NEW_RELIC_ID":123123,
"GEM_PATH":"vendor/bundle/ruby/1.9.1"}`

func TestListConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/apps/foo/config_vars" {
			t.Errorf("invalid URI %q", r.RequestURI)
		}
		header := w.Header()
		header.Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(JSONConfig))
	}))
	defer server.Close()

	client := NewHerokuClient("foobar")
	client.ApiEndpoint = server.URL

	app := types.Application{Name: "foo"}
	config, err := client.ListConfig(&app)
	if err != nil {
		t.Fatalf("expected success: %s", err)
	}
	if len(*config) != 15 {
		t.Fatalf("expected config to have 15 items was %d", len(*config))
	}
}

const JSONAuthError = `{
  "id":"unauthorized",
  "error":"Invalid credentials provided."
}`

func TestErrorApiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write([]byte(JSONAuthError))
	}))
	defer server.Close()

	client := NewHerokuClient("foobar")
	client.ApiEndpoint = server.URL

	_, err := client.ListApplications()

	switch err.(type) {
	case ApiError:
	case nil:
		t.Fatalf("expected an error")
	default:
		t.Fatalf("unexpected error %v", err)
	}

	if err.Error() != "ApiError(unauthorized): Invalid credentials provided." {
		t.Errorf("unexpected message %s", err.Error())
	}
}

func TestErrorUnknownError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Content-Type", "application/json")
		w.WriteHeader(402)
		w.Write([]byte("Payment required to continue."))
	}))

	client := NewHerokuClient("foobar")
	client.ApiEndpoint = server.URL

	_, err := client.ListApplications()

	switch err.(type) {
	case UnknownError:
	case nil:
		t.Fatalf("expected an error")
	default:
		t.Fatalf("unexpected error %v", err)
	}

	if err.Error() != "UnknownError: Payment required to continue." {
		t.Errorf("unexpected message %s", err.Error())
	}

}
