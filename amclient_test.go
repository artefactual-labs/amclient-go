package amclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gotest.tools/v3/assert"
)

var (
	mux    *http.ServeMux
	ctx    = context.TODO()
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	url, _ := url.Parse(server.URL)
	client = NewClient(nil, url.String(), "", "")
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	t.Helper()
	assert.Equal(t, r.Method, expected)
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, err := client.NewRequest(ctx, "GET", "/", nil)
	assert.NilError(t, err)
	body := new(foo)
	_, err = client.Do(context.Background(), req, body)
	assert.NilError(t, err)

	assert.DeepEqual(t, body, &foo{"a"})
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, err := client.NewRequest(ctx, "GET", "/", nil)
	assert.NilError(t, err)
	_, err = client.Do(context.Background(), req, nil)
	assert.Assert(t, err != nil)
}

func TestCustomUserAgent(t *testing.T) {
	c, err := New(nil, "http://127.0.0.1", "", "", SetUserAgent("testing"))
	assert.NilError(t, err)

	expected := fmt.Sprintf("%s+%s", "testing", userAgent)
	assert.Equal(t, c.UserAgent, expected)
}

func TestNewRequest(t *testing.T) {
	var (
		baseURL  = "http://127.0.0.1"
		user     = "Us3r"
		password = "Pa33w0rd"
	)

	c, err := New(nil, baseURL, user, password)
	assert.NilError(t, err)

	inURL, outURL := "/foo", baseURL+"/foo"
	inBody := &TransferStartRequest{Name: "My transfer", Type: "standard"}
	req, err := c.NewRequest(context.Background(), "GET", inURL, inBody)
	assert.NilError(t, err)

	// Test that relative URL was expanded.
	assert.Equal(t, req.URL.String(), outURL)

	// Test that default user-agent is attached to the request.
	assert.Equal(t, req.Header.Get("User-Agent"), c.UserAgent)

	// Test that the Authorization header is included.
	assert.Equal(t, req.Header.Get("Authorization"), fmt.Sprintf("ApiKey %s:%s", user, password))
}

func TestNewRequestJSON(t *testing.T) {
	var (
		baseURL  = "http://127.0.0.1"
		user     = "Us3r"
		password = "Pa33w0rd"
	)

	c, err := New(nil, baseURL, user, password)
	assert.NilError(t, err)

	inBody := struct {
		Test string `json:"string"`
	}{Test: "foobar"}
	req, err := c.NewRequestJSON(context.Background(), "GET", "/foo", inBody)
	assert.NilError(t, err)

	// Test that the Authorization header is included.
	assert.Equal(t, req.Header.Get("Content-Type"), mediaTypeJSON)

	got, err := io.ReadAll(req.Body)
	assert.NilError(t, err)
	want := []byte(`{"string":"foobar"}
`)
	defer req.Body.Close()
	assert.DeepEqual(t, got, want)
}

func TestAddOption(t *testing.T) {
	var (
		rawURL = "http://127.0.0.1:12345"
		err    error
	)

	// It can set a query parameter.
	rawURL, err = addOption(rawURL, "param1", "1")
	assert.NilError(t, err)
	assert.Equal(t, rawURL, "http://127.0.0.1:12345?param1=1")

	// It can overwrite a query parameter.
	rawURL, err = addOption(rawURL, "param1", "2")
	assert.NilError(t, err)
	assert.Equal(t, rawURL, "http://127.0.0.1:12345?param1=2")

	// It can append an additional query parameter.
	rawURL, err = addOption(rawURL, "param2", "a")
	assert.NilError(t, err)
	assert.Equal(t, rawURL, "http://127.0.0.1:12345?param1=2&param2=a")

	// It can set an empty string as a query parameter.
	rawURL, err = addOption(rawURL, "param3", "")
	assert.NilError(t, err)
	assert.Equal(t, rawURL, "http://127.0.0.1:12345?param1=2&param2=a&param3=")
}
