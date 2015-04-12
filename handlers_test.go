package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"
	"math/rand"
	"time"
	"strings"
)

func TestReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://example.com/todos", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoIndex(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Status code incorrect")
	}
}

func TestReturnsJson(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://example.com/todos", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	rand := random(1, 100)
	todoName := "Write more tests " + strconv.Itoa(rand)
	RepoCreateTodo(Todo{Name: todoName})

	TodoIndex(recorder, req)

	expectedJson := `"name":"` + todoName + `"`

	result := recorder.Body.String()

	if strings.Contains(result, expectedJson) != true {
		t.Errorf("json format incorrect: Actual %s, Expected: %s", result, expectedJson)
	}
}

// helpers
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
