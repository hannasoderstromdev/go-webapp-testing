package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var tests = []struct {
		name string
		url string
		expectedStatusCode int
	} {
		{ "home", "/", http.StatusOK },
		{ "404", "/fake", http.StatusNotFound },
	}

	var app application
	routes := app.routes()

	// create test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	// set template dir relative to where the test runs from
	pathToTemplates = "./../../templates/"

	// loop through test data
	for _, e := range tests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("%s: expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}