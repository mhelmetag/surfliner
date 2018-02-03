package surfliner

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type payload struct {
	Data []Area
}

type Area struct {
	ID         int    `json:"id"`
	SurflineID string `json:"surfline_id"`
	Name       string `json:"name"`
}

func DefaultClient() (*Client, error) {
	url, err := url.Parse("http://localhost:4000")
	if err != nil {
		return nil, err
	}

	userAgent := "SurflineR Client"
	httpClient := http.DefaultClient

	return &Client{BaseURL: url, UserAgent: userAgent, httpClient: httpClient}, err
}

func (c *Client) ListAreas() ([]Area, error) {
	rel := &url.URL{Path: "/api/areas"}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p payload
	err = json.NewDecoder(resp.Body).Decode(&p)
	return p.Data, err
}

func (c *Client) get(u *url.URL) (*http.Response, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
