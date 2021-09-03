package selector

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeOut = errors.New("timeout exceeded")

func Racer(a, b string) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-countdown(10 * time.Second):
		return "", ErrTimeOut
	}
}

func countdown(delay time.Duration) chan struct{} {
	ch := make(chan struct{})
	go func(ch chan struct{}, delay time.Duration) {
		time.Sleep(delay)
		close(ch)
	}(ch, delay)

	return ch
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func(ch chan struct{}, url string) {
		http.Get(url)
		close(ch)
	}(ch, url)
	return ch
}
