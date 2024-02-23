package concurrency

import (
	"reflect"
	"testing"
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
