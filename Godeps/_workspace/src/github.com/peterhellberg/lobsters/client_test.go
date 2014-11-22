package lobsters

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	assert.Equal(t, "https://lobste.rs/", c.BaseURL.String())
	assert.Equal(t, "lobsters.go/0.0.1", c.UserAgent)
}

func TestNewRequest(t *testing.T) {
	r, err := NewClient(nil).NewRequest("s/jyloq8.json")

	assert.Nil(t, err)
	assert.Equal(t, "https://lobste.rs/s/jyloq8.json", r.URL.String())
}

func TestDoWithInvalidJSON(t *testing.T) {
	ts, c := testServerAndClientByFixture("invalid")
	defer ts.Close()

	r, _ := c.NewRequest("invalid.json")
	resp, err := c.Do(r, struct{}{})

	assert.Nil(t, resp)
	assert.EqualError(t, err, "error reading response from GET /invalid.json: "+
		"invalid character 'N' looking for beginning of value")
}

func testServerAndClientByFixture(fn string) (*httptest.Server, *Client) {
	body, _ := ioutil.ReadFile("_fixtures/" + fn + ".json")

	ts := testServer(body)

	c := NewClient(nil)
	c.BaseURL, _ = url.Parse(ts.URL)

	return ts, c
}

func testServer(body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}))
}
