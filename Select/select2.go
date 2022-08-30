package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got, err := Racer(slowURL, fastURL)

	if err != nil {
		t.Fatalf("did not expect an error but got one = %v", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(15 * time.Second)
		http.Get(url)
		close(ch)
	}()
	return ch
}
