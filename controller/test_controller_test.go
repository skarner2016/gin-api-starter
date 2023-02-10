package controller_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestTest(t *testing.T) {
	uri := "/test"

	param := url.Values{
		"id":   {"1"},
		"name": {"abc"},
	}

	uri = fmt.Sprintf("%s?%s", uri, param.Encode())

	// w := SendHttpRequest(http.MethodGet, uri, nil)
	w := SendHttpRequest(http.MethodGet, uri, nil)
	fmt.Println("uri:", uri)

	fmt.Println(w.Code, w.Body.String())
}

func TestGetUser(t *testing.T) {
	// uri := "/test/user?id=12&name=abc"
	// uri := "/test/user?id=10000&name=abc"
	// uri := "/test/user?id=10000"
	uri := "/test/user"

	param := url.Values{
		"id":   {"1"},
		"name": {"abc"},
	}

	uri = fmt.Sprintf("%s?%s", uri, param.Encode())

	// w := SendHttpRequest(http.MethodGet, uri, nil)
	w := SendHttpRequest(http.MethodGet, uri, nil)
	fmt.Println("uri:", uri)

	fmt.Println(w.Code, w.Body.String())
}

func TestPostUser(t *testing.T) {
	// uri := "/test/user?id=12&name=abc"
	// uri := "/test/user?id=10000&name=abc"
	// uri := "/test/user?id=10000"
	uri := "/test/user"

	param := url.Values{
		"id":   {"1"},
		"name": {"abc"},
	}

	w := SendHttpRequest(http.MethodPost, uri, param)
	fmt.Println("uri:", uri)

	fmt.Println(w.Code, w.Body.String())
}
