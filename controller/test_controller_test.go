package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"skarner2016/gin-api-starter/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	router := router.Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	fmt.Println(w.Body.String())
}

func TestGetUser(t *testing.T) {
	// uri := "/test/user?id=12&name=abc"
	// uri := "/test/user?id=10000&name=abc"
	// uri := "/test/user?id=10000"
	uri := "/test/user"

	param := url.Values{
		"id":   {"10000"},
		"name": {"abc"},
	}

	uri = fmt.Sprintf("%s?%s", uri, param.Encode())

	// w := SendHttpRequest(http.MethodGet, uri, nil)
	w := SendHttpRequest(http.MethodGet, uri, nil)
	fmt.Println("uri:", uri)

	fmt.Println(w.Code, w.Body.String())
}
