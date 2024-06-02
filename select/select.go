package racer

import (
	"net/http"
	"time"
)

// Race two websites, and return the URL of the website that responded first (timeout after 10s)
func WebsiteRacer(a, b string) (winner string, err error) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	winner = a
	if bDuration < aDuration {
		winner = b
	}
	return winner, err
}
