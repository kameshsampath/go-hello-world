package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGreeterAnonymous(t *testing.T) {
	greeting := Greeting{
		Name:    "Kamesh",
		Message: "Hello,Anonymous",
	}
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, sayHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var got Greeting
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equalf(t, greeting.Message, got.Message, "Expecting %s but got %s", greeting.Message, got.Message)
	}
}

func TestGreetNamed(t *testing.T) {
	greeting := Greeting{
		Name:    "Kamesh",
		Message: "Hello,Kamesh",
	}
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?name=Kamesh", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, sayHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var got Greeting
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equalf(t, greeting.Message, got.Message, "Expecting %s but got %s", greeting.Message, got.Message)
	}
}

func TestGreetNamedPrefix(t *testing.T) {
	greeting := Greeting{
		Name:    "Kamesh",
		Message: "Hi,Kamesh",
	}
	// Setup
	t.Setenv("GREETING_PREFIX", "Hi")
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?name=Kamesh", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, sayHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var got Greeting
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equalf(t, greeting.Message, got.Message, "Expecting %s but got %s", greeting.Message, got.Message)
	}
}
