package surfliner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client is the SurflineR HTTP Client.
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type payloadA struct {
	Data []Area `json:"data"`
}

// Area is the first tier in the Surfline spot hierarchy.
type Area struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type payloadR struct {
	Data []Region `json:"data"`
}

// Region is the second tier in the Surfline spot hierarchy.
type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type payloadSR struct {
	Data []SubRegion `json:"data"`
}

// SubRegion is the third tier in the Surfline spot hierarchy.
type SubRegion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DefaultClient returns a default configured SurflineR Client.
func DefaultClient() (*Client, error) {
	url, err := url.Parse("http://localhost:4000")
	if err != nil {
		return nil, err
	}

	userAgent := "SurflineR Client"
	httpClient := http.DefaultClient
	client := Client{BaseURL: url, UserAgent: userAgent, httpClient: httpClient}

	return &client, err
}

// ListAreas returns all Surfline Areas.
func (c *Client) ListAreas() ([]Area, error) {
	rel := &url.URL{Path: "/api/areas"}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p payloadA
	err = json.NewDecoder(resp.Body).Decode(&p)
	return p.Data, err
}

// ListRegions returns all Surfline Regions for an Area.
func (c *Client) ListRegions(areaID string) ([]Region, error) {
	path := fmt.Sprintf("/api/areas/%s/regions", areaID)
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p payloadR
	err = json.NewDecoder(resp.Body).Decode(&p)
	return p.Data, err
}

// ListSubRegions returns all Surfline SubRegions for a Region.
func (c *Client) ListSubRegions(areaID string, regionID string) ([]SubRegion, error) {
	path := fmt.Sprintf("/api/areas/%s/regions/%s/subregions", areaID, regionID)
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p payloadSR
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
