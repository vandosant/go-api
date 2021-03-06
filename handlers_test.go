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

func TestIndexReturnsText(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	Index(recorder, req)

	expected := "Welcome!"

	result := recorder.Body.String()

	if strings.Contains(result, expected) != true {
		t.Errorf("json format incorrect: Actual %s, Expected: %s", result, expected)
	}
}

func TestTodosReturns200(t *testing.T) {
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

func TestTodoReturnsJson(t *testing.T) {
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

func TestTodoCreateReturns201(t *testing.T) {
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

func TestTodoCreateSavesJSON(t *testing.T) {
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

func TestTodoShowReturnsJSON(t *testing.T) {
	recorder := httptest.NewRecorder()

	rand := random(1, 100)
	todoName := `Write more specs ` + strconv.Itoa(rand)
	requestBody := `{"name":"`+ todoName + `"}`

	req1, err := http.NewRequest("POST", "http://example.com/todos", strings.NewReader(requestBody))
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoCreate(recorder, req1)
	if recorder.Code != http.StatusCreated {
		t.Errorf("Status code incorrect.")
	}

	body := recorder.Body.String()
	id := strings.Split(body, ":")
	idValue := strings.Split(id[1], ",")[0]

	req2, err := http.NewRequest("GET", "http://example.com/todos/"+idValue, nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	TodoShow(recorder, req2)

	result := recorder.Body.String()
	expected := `"name":"`+ todoName + `"`

	if strings.Contains(result, expected) != true {
		t.Errorf("json format incorrect: Actual %s, Expected: %s", result, expected)
	}
}

// helpers
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
