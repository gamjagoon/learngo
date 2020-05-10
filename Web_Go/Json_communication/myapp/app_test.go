package myapp

//Go 는 _test 밑줄을 붙이면 자동으로 test 로 작동하게 된다

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/bar", nil)

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestIndexPathHandler_Name(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/bar?name=tucker", nil)

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello tucker", string(data))
}

func TestIndexPathHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/foo", nil)

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadGateway, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Bad Request: EOF", string(data))
}

func TestIndexPathHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"fisrt_name":"minjae","last_name":"kim","email":"rlaalswo201@gmail.com"}`))

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
}
