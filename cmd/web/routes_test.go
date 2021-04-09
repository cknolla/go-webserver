package main

import (
	"github.com/cknolla/go-webserver/internal/config"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch mux.(type) {
	case http.Handler:
	default:
		t.Error("not a handler")
	}
}
