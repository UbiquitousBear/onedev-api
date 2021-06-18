package onedev

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", http.StripPrefix("/", mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		http.Error(w, "path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.

	baseEndpoint, _ := url.Parse(server.URL + "/")
	client, _ = NewClient(baseEndpoint.String())

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestDo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	body := new(foo)
	ctx := context.Background()
	client.Do(ctx, req, body)

	want := &foo{"a"}
	if !cmp.Equal(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_nilContext(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(nil, req, nil)

	if !errors.Is(err, errors.New("context must be non-nil")) {
		t.Errorf("Expected context must be non-nil error")
	}
}

func TestDo_httpError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	ctx := context.Background()
	resp, err := client.Do(ctx, req, nil)

	if err == nil {
		t.Fatal("Expected HTTP 400 error, got no error.")
	}
	if resp.StatusCode != 400 {
		t.Errorf("Expected HTTP 400 error, got %d status code.", resp.StatusCode)
	}
}