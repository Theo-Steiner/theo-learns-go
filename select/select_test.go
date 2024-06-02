package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestWebsiteRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got, err := WebsiteRacer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("received error %q", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respon within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(25 * time.Millisecond)
		defer slowServer.Close()

		_, err := ConfigurableWebsiteRacer(slowServer.URL, slowServer.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
