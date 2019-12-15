package httpselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL
	t.Run("a is faster", func(t *testing.T) {
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})
	t.Run("b is faster", func(t *testing.T) {
		got := Racer(fastURL, slowURL)

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
