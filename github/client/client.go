package client

import (
	"net/http"
	"time"
)

// Create a structurre to pass client information around with
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

const HostURL string = "https://api.github.com/"

func NewClient(token *string) (*Client, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	return &c, nil
}
