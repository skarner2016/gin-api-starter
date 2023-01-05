package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"skarner2016/gin-api-starter/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {

	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	fmt.Println(w.Body.String())
}

func TestGetUser(t *testing.T) {
	router := router.SetupRouter()

	// uri := "/test/user?id=12&name=abc"
	uri := "/test/user?id=10000&name=abc"
	// uri := "/test/user?id=10000"
	// uri := "/test/user"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	fmt.Println(w.Body.String())
}
