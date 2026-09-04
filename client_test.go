package swervpay

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	client *SwervpayClient
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewSwervpayClient(&SwervpayClientOption{
		BusinessID: "",
		SecretKey:  "",
		BaseURL:    "",
	})
	baseURL := server.URL
	if baseURL[len(baseURL)-1] != '/' {
		baseURL += "/"
	}
	parsedURL, _ := url.Parse(baseURL)
	client.BaseURL = parsedURL
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func TestAuthenticate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		username, password, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "business", username)
		assert.Equal(t, "secret", password)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"token_001","token":{"type":"bearer","issued_at":1}}`))
	})

	client.Config.BusinessID = "business"
	client.Config.SecretKey = "secret"
	response, err := client.Auth.Authenticate(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "token_001", response.AccessToken)
	assert.Equal(t, "token_001", client.AccessToken)
}
