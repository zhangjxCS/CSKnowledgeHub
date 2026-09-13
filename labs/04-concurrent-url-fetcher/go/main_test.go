package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestFetchOneSuccess(t *testing.T) {
	const body = "hello, fetcher"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()

	result := FetchOne(context.Background(), server.Client(), server.URL, time.Second)

	if result.URL != server.URL {
		t.Fatalf("URL = %q, want %q", result.URL, server.URL)
	}
	if result.StatusCode != http.StatusCreated {
		t.Errorf("StatusCode = %d, want %d", result.StatusCode, http.StatusCreated)
	}
	if result.BodySize != len(body) {
		t.Errorf("BodySize = %d, want %d", result.BodySize, len(body))
	}
	if result.Err != nil {
		t.Errorf("Err = %v, want nil", result.Err)
	}
}

func TestFetchAllPreservesOrderAndDuplicates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/slow" {
			time.Sleep(50 * time.Millisecond)
		}
		_, _ = w.Write([]byte(r.URL.Path))
	}))
	defer server.Close()

	urls := []string{
		server.URL + "/slow",
		server.URL + "/fast",
		server.URL + "/fast",
	}

	results, err := FetchAll(context.Background(), server.Client(), urls, 3, time.Second)
	if err != nil {
		t.Fatalf("FetchAll() error = %v", err)
	}
	if len(results) != len(urls) {
		t.Fatalf("len(results) = %d, want %d", len(results), len(urls))
	}

	for i, wantURL := range urls {
		if results[i].URL != wantURL {
			t.Errorf("results[%d].URL = %q, want %q", i, results[i].URL, wantURL)
		}
		if results[i].Err != nil {
			t.Errorf("results[%d].Err = %v, want nil", i, results[i].Err)
		}
	}
}

func TestFetchAllRespectsConcurrencyLimit(t *testing.T) {
	const (
		urlCount = 8
		limit    = 3
	)

	var active atomic.Int32
	var maximum atomic.Int32
	entered := make(chan struct{}, urlCount)
	release := make(chan struct{})
	var releaseOnce sync.Once
	releaseAll := func() {
		releaseOnce.Do(func() { close(release) })
	}
	defer releaseAll()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		current := active.Add(1)
		defer active.Add(-1)

		for {
			observed := maximum.Load()
			if current <= observed || maximum.CompareAndSwap(observed, current) {
				break
			}
		}

		entered <- struct{}{}
		select {
		case <-release:
		case <-r.Context().Done():
		}
	}))
	defer server.Close()

	urls := make([]string, urlCount)
	for i := range urls {
		urls[i] = server.URL
	}

	type fetchOutcome struct {
		results []FetchResult
		err     error
	}
	done := make(chan fetchOutcome, 1)
	go func() {
		results, err := FetchAll(context.Background(), server.Client(), urls, limit, time.Second)
		done <- fetchOutcome{results: results, err: err}
	}()

	for i := 0; i < limit; i++ {
		select {
		case <-entered:
		case <-time.After(time.Second):
			t.Fatalf("only %d requests started; expected %d", i, limit)
		}
	}

	select {
	case <-entered:
		t.Fatalf("more than %d requests were active at once", limit)
	case <-time.After(30 * time.Millisecond):
	}

	releaseAll()
	outcome := <-done
	if outcome.err != nil {
		t.Fatalf("FetchAll() error = %v", outcome.err)
	}
	if len(outcome.results) != urlCount {
		t.Fatalf("len(results) = %d, want %d", len(outcome.results), urlCount)
	}
	if got := maximum.Load(); got != limit {
		t.Errorf("maximum active requests = %d, want %d", got, limit)
	}
}

func TestFetchAllTimeoutDoesNotStopOtherURLs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/slow" {
			select {
			case <-time.After(time.Second):
			case <-r.Context().Done():
				return
			}
		}
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	urls := []string{server.URL + "/slow", server.URL + "/fast"}
	results, err := FetchAll(context.Background(), server.Client(), urls, 2, 30*time.Millisecond)
	if err != nil {
		t.Fatalf("FetchAll() error = %v", err)
	}

	if !errors.Is(results[0].Err, context.DeadlineExceeded) {
		t.Errorf("slow request error = %v, want context deadline exceeded", results[0].Err)
	}
	if results[1].Err != nil {
		t.Errorf("fast request error = %v, want nil", results[1].Err)
	}
	if results[1].StatusCode != http.StatusOK {
		t.Errorf("fast request status = %d, want %d", results[1].StatusCode, http.StatusOK)
	}
}

func TestFetchAllCancellationReturnsOneResultPerURL(t *testing.T) {
	started := make(chan struct{}, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started <- struct{}{}
		<-r.Context().Done()
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	urls := []string{server.URL, server.URL, server.URL, server.URL}

	type fetchOutcome struct {
		results []FetchResult
		err     error
	}
	done := make(chan fetchOutcome, 1)
	go func() {
		results, err := FetchAll(ctx, server.Client(), urls, 1, time.Second)
		done <- fetchOutcome{results: results, err: err}
	}()

	select {
	case <-started:
		cancel()
	case <-time.After(time.Second):
		t.Fatal("request did not start")
	}

	select {
	case outcome := <-done:
		if outcome.err != nil {
			t.Fatalf("FetchAll() error = %v", outcome.err)
		}
		if len(outcome.results) != len(urls) {
			t.Fatalf("len(results) = %d, want %d", len(outcome.results), len(urls))
		}
		for i, result := range outcome.results {
			if result.URL != urls[i] {
				t.Errorf("results[%d].URL = %q, want %q", i, result.URL, urls[i])
			}
			if !errors.Is(result.Err, context.Canceled) {
				t.Errorf("results[%d].Err = %v, want context canceled", i, result.Err)
			}
		}
	case <-time.After(time.Second):
		t.Fatal("FetchAll did not return promptly after cancellation")
	}
}

func TestFetchAllReportsFailuresIndependently(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/missing" {
			w.WriteHeader(http.StatusNotFound)
		}
		_, _ = w.Write([]byte("response"))
	}))
	defer server.Close()

	urls := []string{
		server.URL + "/ok",
		"://malformed",
		server.URL + "/missing",
	}
	results, err := FetchAll(context.Background(), server.Client(), urls, 3, time.Second)
	if err != nil {
		t.Fatalf("FetchAll() error = %v", err)
	}

	if results[0].Err != nil || results[0].StatusCode != http.StatusOK {
		t.Errorf("successful result = %+v", results[0])
	}
	if results[1].Err == nil {
		t.Error("malformed URL error = nil, want an error")
	}
	if results[2].Err != nil {
		t.Errorf("non-2xx error = %v, want nil", results[2].Err)
	}
	if results[2].StatusCode != http.StatusNotFound {
		t.Errorf("non-2xx status = %d, want %d", results[2].StatusCode, http.StatusNotFound)
	}
}

func TestFetchAllValidatesArguments(t *testing.T) {
	tooManyURLs := make([]string, 101)
	for i := range tooManyURLs {
		tooManyURLs[i] = fmt.Sprintf("http://example.test/%d", i)
	}

	tests := []struct {
		name        string
		client      *http.Client
		urls        []string
		concurrency int
		timeout     time.Duration
	}{
		{name: "too many URLs", client: http.DefaultClient, urls: tooManyURLs, concurrency: 1, timeout: time.Second},
		{name: "zero concurrency", client: http.DefaultClient, concurrency: 0, timeout: time.Second},
		{name: "negative concurrency", client: http.DefaultClient, concurrency: -1, timeout: time.Second},
		{name: "zero timeout", client: http.DefaultClient, concurrency: 1, timeout: 0},
		{name: "nil client", client: nil, concurrency: 1, timeout: time.Second},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := FetchAll(
				context.Background(),
				test.client,
				test.urls,
				test.concurrency,
				test.timeout,
			)
			if err == nil {
				t.Fatal("FetchAll() error = nil, want validation error")
			}
		})
	}
}

type trackingBody struct {
	reader *strings.Reader
	closed atomic.Bool
}

func (b *trackingBody) Read(p []byte) (int, error) {
	return b.reader.Read(p)
}

func (b *trackingBody) Close() error {
	b.closed.Store(true)
	return nil
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestFetchOneReadsAndClosesResponseBody(t *testing.T) {
	body := &trackingBody{reader: strings.NewReader("chunked body")}
	client := &http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode:    http.StatusOK,
				Body:          body,
				ContentLength: -1,
				Header:        make(http.Header),
			}, nil
		}),
	}

	result := FetchOne(context.Background(), client, "http://example.test", time.Second)

	if result.Err != nil {
		t.Fatalf("FetchOne() error = %v", result.Err)
	}
	if result.BodySize != len("chunked body") {
		t.Errorf("BodySize = %d, want %d", result.BodySize, len("chunked body"))
	}
	if !body.closed.Load() {
		t.Error("response body was not closed")
	}
}
