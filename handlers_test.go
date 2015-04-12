package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://example.com/tasks", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoIndex(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Status code incorrect")
	}
}
