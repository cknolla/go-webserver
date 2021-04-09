package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	handler := NoSurf(&myH)

	switch handler.(type) {
	case http.Handler:
	default:
		t.Error("type is not http.Handler")
	}
}