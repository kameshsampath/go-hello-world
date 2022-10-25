package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, sayHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		got := rec.Body.String()
		assert.Equalf(t, greeting, got, "Expecting %s but got %s", greeting, got)
	}
}
