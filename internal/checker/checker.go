package checker

import (
	"context"
	"net/http"
	"time"
)

const waitingTime = 3 * time.Second

func ExistUrl(url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), waitingTime)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest
}
