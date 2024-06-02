package racer

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// Race two websites, and return the URL of the website that responded first (timeout after 10s)
func WebsiteRacer(a, b string) (winner string, err error) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a, err
	}
	return b, err
}
