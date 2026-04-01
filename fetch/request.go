package fetch

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Request ...
type Request struct {
	// Body is the request body as a string.
	Body string
	// Method is the HTTP method of the request.
	Method string
	// Redirect is the URL to redirect to, if any.
	Redirect string
	// Header is the request headers.
	Header http.Header
	// URL is the URL of the request.
	URL *url.URL
	// RemoteAddr is the remote address of the request.
	RemoteAddr string
}

// ParseRequestURL ...
func ParseRequestURL(rawURL string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url '%s' is not valid, %w", rawURL, err)
	}

	switch u.Scheme {
	case "http", "https":
	case "": // then scheme is empty, it's a local request
		if !strings.HasPrefix(u.Path, "/") {
			return nil, fmt.Errorf("unsupported relative path %s", u.Path)
		}
	default:
		return nil, fmt.Errorf("unsupported scheme %s", u.Scheme)
	}

	return u, nil
}
