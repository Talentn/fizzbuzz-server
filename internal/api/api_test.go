package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/talentn/fizzbuzz-server/internal/stats"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	ts := httptest.NewServer(NewHandler(stats.New()))
	t.Cleanup(ts.Close)
	return ts
}

func getJSON(t *testing.T, url string, wantStatus int) map [string]any {
	t.Helper()

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != wantStatus {
		t.Fatalf("GET %s: %d, want %d", url, resp.StatusCode, wantStatus)
	}

	var body map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	return body
}

func TestFizzbuzzEndpoint(t *testing.T) {
	ts := newTestServer(t)

	body := getJSON(t, ts.URL+"/fizzbuzz?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz", http.StatusOK)

	result, ok := body["result"].([]any)
	if !ok {
		t.Fatalf("body['result'] is not an array: %v", body["result"])
	}
	if len(result) != 15 {
		t.Fatalf("body['result'] has %d elements, want 15", len(result))
	}
	if result[2] != "Fizz" || result[4] != "Buzz" || result[14] != "FizzBuzz" {
		t.Fatalf("body['result'] has unexpected values: %v", result)
	}
}

func TestFizzBuzzRejectsBadInput(t *testing.T) {
	ts := newTestServer(t)
	badQueries := []string{
		"int1=3",
		"int1=abc&int2=5&limit=10&str1=a&str2=b",
		"int1=0&int2=5&limit=10&str1=a&str2=b",
		"int1=3&int2=5&limit=999999999&str1=a&str2=b",
	}
	for _, q := range badQueries {
		body := getJSON(t, ts.URL+"/fizzbuzz?"+q, http.StatusBadRequest)
		if _, ok := body["errors"]; !ok {
			t.Errorf("query %q: response has no errors field: %v", q, body)
		}
	}
}
func TestStatsEndpoint(t *testing.T) {
	ts := newTestServer(t)
	popular := "/fizzbuzz?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz"
	other := "/fizzbuzz?int1=2&int2=7&limit=10&str1=a&str2=b"
	getJSON(t, ts.URL+popular, http.StatusOK)
	getJSON(t, ts.URL+popular, http.StatusOK)
	getJSON(t, ts.URL+other, http.StatusOK)
	body := getJSON(t, ts.URL+"/stats", http.StatusOK)
	if hits, _ := body["hits"].(float64); hits != 2 {
		t.Errorf("hits = %v, want 2", body["hits"])
	}
}
