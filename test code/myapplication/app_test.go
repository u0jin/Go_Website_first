package myapplication


import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"strings"
	"github.com/stretchr/testify/assert"

)

func TestIndexPathHandler(t *testing.T){ 
	assert := assert.New(t)
	
	res := httptest.NewRecorder() 
	req := httptest.NewRequest("GET", "/", nil)
	
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)

	
	assert.Equal("Hello Ujin", string(data))

}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}


func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=ujin", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello ujin!", string(data))
}


func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)

}


func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"youjin", "last_name":"kim", "email":"youjin@studnetpartner.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code) 

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("youjin", user.FirstName)
	assert.Equal("kim", user.LastName)

}