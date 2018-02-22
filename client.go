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

// Place can either be an Area, Region or SubRegion (any of the tiers in Surfline's Place hierarchy).
type Place struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// APIError is an error that can be returned by the Surfline Regions API.
type APIError struct {
	msg        string
	StatusCode int
}

func (e APIError) Error() string {
	return e.msg
}

type dPs struct {
	Data []Place `json:"data"`
}

type dP struct {
	Data Place `json:"data"`
}

// DefaultClient returns a default configured SurflineR Client.
func DefaultClient() (*Client, error) {
	url, err := url.Parse("http://surfliner.maxworld.tech")
	if err != nil {
		return nil, err
	}

	userAgent := "SurflineR Client"
	httpClient := http.DefaultClient
	client := Client{BaseURL: url, UserAgent: userAgent, httpClient: httpClient}

	return &client, err
}

// Areas returns all Surfline Areas.
func (c *Client) Areas() ([]Place, error) {
	rel := &url.URL{Path: "/api/areas"}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var p dPs
	err = json.NewDecoder(resp.Body).Decode(&p)
	return p.Data, err
}

// Regions returns all Surfline Regions for an Area.
func (c *Client) Regions(areaID string) ([]Place, error) {
	path := fmt.Sprintf("/api/areas/%s/regions", areaID)
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = handleAPIErrors(resp.StatusCode)
	if err != nil {
		return nil, err
	}

	var p dPs
	err = json.NewDecoder(resp.Body).Decode(&p)
	return p.Data, err
}

// SubRegions returns all Surfline SubRegions for a Region.
func (c *Client) SubRegions(areaID string, regionID string) ([]Place, error) {
	path := fmt.Sprintf("/api/areas/%s/regions/%s/subregions", areaID, regionID)
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = handleAPIErrors(resp.StatusCode)
	if err != nil {
		return nil, err
	}

	var p dPs
	err = json.NewDecoder(resp.Body).Decode(&p)
	return p.Data, err
}

// SubRegion returns a Surfline SubRegion.
func (c *Client) SubRegion(areaID string, regionID string, subRegionID string) (Place, error) {
	path := fmt.Sprintf("/api/areas/%s/regions/%s/subregions/%s", areaID, regionID, subRegionID)
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	resp, err := c.get(u)
	if err != nil {
		return Place{}, err
	}
	defer resp.Body.Close()

	err = handleAPIErrors(resp.StatusCode)
	if err != nil {
		return Place{}, err
	}

	var p dP
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

func handleAPIErrors(code int) error {
	var msg string

	if code == http.StatusOK {
		return nil
	} else if code == http.StatusNotFound {
		msg = "the specified place could not be found"
	} else {
		msg = "an error occured while making a request"
	}

	return APIError{msg, code}
}
