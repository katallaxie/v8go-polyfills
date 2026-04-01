package fetch_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/katallaxie/v8go-polyfills/fetch"
	"github.com/stretchr/testify/assert"
)

func TestNewResponseFromHTTP(t *testing.T) {
	tests := []struct {
		name    string
		res     *http.Response
		want    *fetch.Response
		wantErr bool
	}{
		{
			name: "200 OK",
			res: &http.Response{
				StatusCode: 200,
				Status:     "200 OK",
				Header:     http.Header{"Content-Type": []string{"text/plain"}},
				Body:       io.NopCloser(strings.NewReader("Hello, World!")),
				Request: &http.Request{
					URL:    &url.URL{Scheme: "http", Host: "example.com", Path: "/"},
					Method: "GET",
					Header: http.Header{
						"Referer": []string{"http://example.com/"},
					},
				},
			},
			want: &fetch.Response{
				Body:       "Hello, World!",
				Header:     http.Header{"Content-Type": []string{"text/plain"}},
				OK:         true,
				Redirected: false,
				Status:     200,
				StatusText: "OK",
				URL:        "http://example.com/",
			},
			wantErr: false,
		},
		{
			name: "404 Not Found",
			res: &http.Response{
				StatusCode: 404,
				Status:     "404 Not Found",
				Header:     http.Header{"Content-Type": []string{"text/plain"}},
				Body:       io.NopCloser(strings.NewReader("Not Found")),
				Request: &http.Request{
					URL:    &url.URL{Scheme: "http", Host: "example.com", Path: "/notfound"},
					Method: "GET",
					Header: http.Header{
						"Referer": []string{"http://example.com/notfound"},
					},
				},
			},
			want: &fetch.Response{
				Body:       "Not Found",
				Header:     http.Header{"Content-Type": []string{"text/plain"}},
				OK:         false,
				Redirected: false,
				Status:     404,
				StatusText: "Not Found",
				URL:        "http://example.com/notfound",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetch.NewResponseFromHTTP(tt.res)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
