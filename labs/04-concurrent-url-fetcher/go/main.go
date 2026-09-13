package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type FetchResult struct {
	URL        string
	StatusCode int
	BodySize   int
	Err        error
}

type job struct {
	index int
	url   string
}

func FetchAll(ctx context.Context, client *http.Client, urls []string, concurrency int, timeout time.Duration) ([]FetchResult, error) {
	if len(urls) > 100 {
		return nil, fmt.Errorf("too many URLs: %d (max 100)", len(urls))
	}
	if concurrency <= 0 {
		return nil, fmt.Errorf("concurrency must be greater than 0")
	}
	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be greater than 0")
	}
	if client == nil {
		return nil, fmt.Errorf("http client cannot be nil")
	}

	results := make([]FetchResult, len(urls))
	jobs := make(chan job, len(urls))
	for url := range urls {
		job := job{index: url, url: urls[url]}
		jobs <- job
	}
	close(jobs)

	var wg sync.WaitGroup
	var workerCount int
	if concurrency > len(urls) {
		workerCount = len(urls)
	} else {
		workerCount = concurrency
	}
	wg.Add(workerCount)
	worker := func() {
		defer wg.Done()
		for j := range jobs {
			if err := ctx.Err(); err != nil {
				results[j.index] = FetchResult{URL: j.url, Err: err}
				continue
			}
			result := FetchOne(ctx, client, j.url, timeout)
			results[j.index] = result
		}
	}

	for range workerCount {
		go worker()
	}

	wg.Wait()
	return results, nil
}

func FetchOne(ctx context.Context, client *http.Client, url string, timeout time.Duration) FetchResult {
	result := FetchResult{URL: url}

	requestCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(requestCtx, http.MethodGet, url, nil)
	if err != nil {
		result.Err = err
		return result
	}

	resp, err := client.Do(req)
	if err != nil {
		result.Err = err
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode

	bodySize, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		result.Err = err
		return result
	}
	result.BodySize = int(bodySize)

	return result
}
