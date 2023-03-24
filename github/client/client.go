package client

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
	"time"
)

// Create a structure to pass client information around with
type Client struct {
	HostURL    string
	Token      string
	HTTPClient *http.Client
	Diags      diag.Diagnostics
}

const HostURL string = "https://api.github.com"

func NewClient(token *string) *Client {

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,            // Base Host URL
		Diags:      diag.Diagnostics{}, // Pointer to Diagnostics, NIL if no error
	}

	return &c
}
