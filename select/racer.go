package racer

import (
	"net/http"
	"time"
)

func Racer(firstUrl, secondUrl string) string {
	startFirstUrl := time.Now()
	http.Get(firstUrl)
	firstUrlDuration := time.Since(startFirstUrl)

	startSecondUrl := time.Now()
	http.Get(secondUrl)
	secondUrlDuration := time.Since(startSecondUrl)
	
	if firstUrlDuration > secondUrlDuration {
		return secondUrl
	}

	return firstUrl
}
