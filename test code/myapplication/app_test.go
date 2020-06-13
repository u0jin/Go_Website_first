package myapplication


import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
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
