package fetch

import (
	"io"
	"net/http"
)

// Response ...
type Response struct {
	// Body is the response body as a string.
	Body string
	// Header is the response headers.
	Header http.Header
	// OK is true if the status code is in the range 200-299.
	OK bool
	// Redirected is true if the response was redirected.
	Redirected bool
	// Status is the HTTP status code.
	Status int32
	// StatusText is the HTTP status text.
	StatusText string
	// URL is the URL of the response.
	URL string
}

// NewResponseFromHTTP ...
func NewResponseFromHTTP(res *http.Response) (*Response, error) {
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:       string(resBody),
		Header:     res.Header,
		OK:         res.StatusCode >= 200 && res.StatusCode < 300,
		Redirected: res.Request.URL.String() != res.Request.Referer(),
		Status:     int32(res.StatusCode),
		StatusText: http.StatusText(res.StatusCode),
		URL:        res.Request.URL.String(),
	}, nil
}
