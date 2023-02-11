package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
    t.Run("it should test time of response.", func(t *testing.T) {
        slowServer := makeDelayedServer(20 * time.Millisecond)
        fastServer := makeDelayedServer(0 * time.Millisecond)

        defer slowServer.Close()
        defer fastServer.Close()

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        want := fastURL
        got, _ := Racer(slowURL, fastURL)

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })

    t.Run("it should return an error if Racer takes longer than 10 seconds.", func(t *testing.T) {
        firstServer := makeDelayedServer(11 * time.Second)
        secondServer := makeDelayedServer(12 * time.Second)

        defer firstServer.Close()
        defer secondServer.Close()

        _, err := Racer(firstServer.URL, secondServer.URL)

        if err == nil {
            t.Errorf("didn't get an error. which is a bad thing.")
        }
    })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}

