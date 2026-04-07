package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}
func TestHelloEndpoint1(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/hello1", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}
func TestHelloEndpoint2(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/hello2", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}
func TestHelloEndpoint3(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/hello3", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}
func TestHelloEndpoint4(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/hello4", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}
