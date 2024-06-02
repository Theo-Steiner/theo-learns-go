package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func SyncWebsiteRacer(a, b string) (winner string, err error) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a, err
	}
	return b, err
}

func ping(url string) chan struct{} {
	// create channel
	// since we will never send to the channel the type being
	// `struct{}` is the smallest available type in memory
	ch := make(chan struct{})
	// concurrently starts goroutine
	go func() {
		// blocks until the response comes back
		http.Get(url)
		// write to channel as soon as response comes back
		close(ch)
	}()
	// immediately returns channel (before the goroutine finishes)
	return ch
}

func ConfigurableWebsiteRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// select takes whichever value comes first
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// Race two websites, and return the URL of the website that responded first (timeout after 10s)
func WebsiteRacer(a, b string) (winner string, err error) {
	return ConfigurableWebsiteRacer(a, b, tenSecondTimeout)
}
