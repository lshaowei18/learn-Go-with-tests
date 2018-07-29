package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Normal race", func(t *testing.T) {
		slowServer := CreateDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := CreateDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		got, _ := Racer(slowServer.URL, fastServer.URL)
		want := fastServer.URL
		if got != want {
			t.Errorf("got %s, want %s.", want, got)
		}
	})

	t.Run("More than 10sec race", func(t *testing.T) {
		slowServer := CreateDelayedServer(11 * time.Second)
		defer slowServer.Close()
		fastServer := CreateDelayedServer(12 * time.Second)
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Error("wanted error, got none.")
		}
	})
}

func CreateDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
