package freshdesk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	subdomain string
	apiKey    string
	baseURL   string

	httpClient *http.Client
}

func NewClient(subdomain, apiKey string) (*Client, error) {
	return &Client{
		apiKey:     apiKey,
		baseURL:    fmt.Sprintf("https://%s.freshdesk.com/api/v2/", subdomain),
		httpClient: retryablehttp.NewClient().StandardClient(),
	}, nil
}

func (c *Client) Tickets() *TicketsClient {
	return &TicketsClient{c}
}

func (c *Client) newRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) do(req *http.Request, out interface{}) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return errors.New(http.StatusText(res.StatusCode))
	}

	if out != nil {
		if err = json.NewDecoder(res.Body).Decode(out); err != nil {
			return err
		}
	}

	return nil
}
