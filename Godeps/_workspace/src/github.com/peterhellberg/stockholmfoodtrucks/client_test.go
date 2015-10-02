package stockholmfoodtrucks

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	for i, tt := range []struct {
		httpClients []*http.Client
	}{
		{nil},
		{[]*http.Client{&http.Client{}}},
	} {
		var c *Client

		if tt.httpClients != nil {
			c = NewClient(tt.httpClients...)
		} else {
			c = NewClient()
		}

		if got, want := c.UserAgent, "stockholmfoodtrucks.go"; got != want {
			t.Fatalf(`[%d] c.UserAgent = %q, want %q`, i, got, want)
		}
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient()

	for i, tt := range []struct {
		path string
	}{
		{""},
		{"/foo"},
	} {
		req, err := c.NewRequest(tt.path)
		if err != nil {
			t.Fatalf(`[%d] %v`, i, err)
		}

		if got, want := req.URL.Path, tt.path; got != want {
			t.Fatalf(`[%d] req.URL.Path = %q, want %q`, i, got, want)
		}
	}
}

func TestDo(t *testing.T) {
	bodyString := "the test body"

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(bodyString))
		},
	))
	defer server.Close()

	c := NewClient(&http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}})

	req, err := c.NewRequest("")
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	res, err := c.Do(req)
	if err != nil {
		t.Fatalf(`%v`, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if got, want := string(body), bodyString; got != want {
		t.Fatalf(`string(body) = %q, want %q`, got, want)
	}
}

func TestNewDocument(t *testing.T) {
	bodyString := `<html><body><div id="foo"><ul><li></li><li class="bar">baz</li></ul></div></body><html>`

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(bodyString))
		},
	))
	defer server.Close()

	c := NewClient(&http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}})

	doc, err := c.NewDocument("")
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if got, want := len(doc.Find("li").Nodes), 2; got != want {
		t.Fatalf(`len(doc.Find("li").Nodes) = %d, want %d`, got, want)
	}
}

func TestEnv(t *testing.T) {
	in, out := "baz", "bar"

	os.Setenv("ENVSTR", out)

	if got := Env("ENVSTR", in); got != out {
		t.Errorf(`String("ENVSTR", "%v") = %v, want %v`, in, got, out)
	}
}

func TestEnvDefault(t *testing.T) {
	in, out := "baz", "baz"

	if got := Env("ENVSTR_DEFAULT", in); got != out {
		t.Errorf(`String("ENVSTR_DEFAULT", "%v") = %v, want %v`, in, got, out)
	}
}
