package concurrency

import (
	"reflect"
	"testing"
	"time"
)

var invalidUrl = "waat://furhurterwe.geds"

func mockWebsiteChecker(url string) bool {
	return url != invalidUrl
}

func TestCheckWebsites(t *testing.T) {
	want := map[string]bool{
		"https://google.com": true,
		"https://aws.com":    true,
		invalidUrl:           false,
	}

	var websites []string
	for k, _ := range want {
		websites = append(websites, k)
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v", want, got)
	}
}

// Benchmarks

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// to run benchmark tests: go test -v -bench=.

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
