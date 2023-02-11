package racer

import (
	"errors"
	"net/http"
	"time"
)
	
func Racer(firstUrl, secondUrl string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	case <-time.After(timeout * time.Second):
	   return "", errors.New("connection timed out for both servers.")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
		
	return ch
}

