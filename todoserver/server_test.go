package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServer(t *testing.T) {

	t.Run("returns Pepper's todo", func(t *testing.T) {
		//New request
		request, _ := http.NewRequest(http.MethodGet, "/users/Pepper", nil)
		response := httptest.NewRecorder()

		UserServer(response, request)

		assertResBody(t, response.Body.String(), "Write Code")
	})

	t.Run("returns Shackleton's todo", func(t *testing.T) {
		//New request
		request, _ := http.NewRequest(http.MethodGet, "/users/Shackleton", nil)
		response := httptest.NewRecorder()

		UserServer(response, request)

		assertResBody(t, response.Body.String(), "Sail")
	})

	t.Run("User that doesn't exist", func(t *testing.T) {
		//New request
		request, _ := http.NewRequest(http.MethodGet, "/users/gkdajgd", nil)
		response := httptest.NewRecorder()

		UserServer(response, request)
		assertStatusCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("URL shorter than /users/", func(t *testing.T) {
		//New request
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		UserServer(response, request)
		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func assertResBody(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("Wrong status code, got %d, want %d", got, want)
	}
}
