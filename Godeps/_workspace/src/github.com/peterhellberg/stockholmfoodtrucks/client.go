package stockholmfoodtrucks

import (
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// A Client communicates with stockholmfoodtrucks.nu
type Client struct {
	// URL is the url for stockholmfoodtrucks.nu
	URL *url.URL

	// User agent used for HTTP requests to stockholmfoodtrucks.nu
	UserAgent string

	// HTTP client used to communicate with stockholmfoodtrucks.nu
	httpClient *http.Client
}

// NewClient returns a new stockholmfoodtrucks client.
func NewClient(httpClients ...*http.Client) *Client {
	cloned := *http.DefaultClient
	httpClient := &cloned

	if len(httpClients) > 0 && httpClients[0] != nil {
		httpClient = httpClients[0]
	}

	return &Client{
		URL: &url.URL{
			Scheme: Env("STOCKHOLM_FOOD_TRUCKS_URL_SCHEME", "http"),
			Host:   Env("STOCKHOLM_FOOD_TRUCKS_URL_HOST", "stockholmfoodtrucks.nu"),
		},
		UserAgent:  Env("STOCKHOLM_FOOD_TRUCKS_USER_AGENT", "stockholmfoodtrucks.go"),
		httpClient: httpClient,
	}
}

// NewRequest creates a new request to stockholmfoodtrucks.nu
func (c *Client) NewRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.URL.String()+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends a request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// Make sure to close the connection after replying to this request
	req.Close = true

	return c.httpClient.Do(req)
}

// NewDocument returns a goquery document based on stockholmfoodtrucks.nu
func (c *Client) NewDocument(path string) (*goquery.Document, error) {
	req, err := c.NewRequest(path)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return goquery.NewDocumentFromResponse(res)
}

// Env returns a string from the ENV, or fallback variable
func Env(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return fallback
}
