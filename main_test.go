package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/takama/router"
)

func TestHandler(t *testing.T) {
	r := router.New()
	r.GET("/", home)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedGreeting := "URL.Path = \"/\""
	testingGreeting := strings.Trim(string(greeting), "\n")
	if testingGreeting != expectedGreeting {
		t.Fatalf("Wrong greeting '%s', expected '%s'", testingGreeting, expectedGreeting)
	}
}
