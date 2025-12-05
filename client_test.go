package swervpay

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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
