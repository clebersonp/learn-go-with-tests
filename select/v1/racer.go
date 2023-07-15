package racer

import (
	"net/http"
	"time"
)

func Racer(urlA string, urlB string) (winner string) {
	durationA := measureResponseTime(urlA)
	durationB := measureResponseTime(urlB)

	if durationA < durationB {
		return urlA
	}

	return urlB
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
