package main

import (
	"testing"

	"github.com/arkarhtethan/golang-web-booking/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Errorf("type is not *chi.mux, but is %T", v)
	}
}
