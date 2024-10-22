package main

import (
	"context"
	"net/http"
	"time"
)

func CheckUrl(url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
