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

func TestReturns201(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "http://example.com/todos/new", strings.NewReader(`{"name":"Write more tests"}`))
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoCreate(recorder, req)
	if recorder.Code != http.StatusCreated {
		t.Errorf("Status code incorrect.")
	}
}

func TestSavesJSON(t *testing.T) {
	recorder := httptest.NewRecorder()

	rand := random(1, 100)
	todoName := `Write more tests ` + strconv.Itoa(rand)
	expectedBody := `{"name":"`+ todoName + `"}`

	req1, err := http.NewRequest("POST", "http://example.com/todos", strings.NewReader(expectedBody))
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoCreate(recorder, req1)
	if recorder.Code != http.StatusCreated {
		t.Errorf("Status code incorrect.")
	}

	req2, err := http.NewRequest("GET", "http://example.com/todos", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoIndex(recorder, req2)

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
