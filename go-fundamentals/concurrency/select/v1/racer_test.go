package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://www.facebook.com"
	fastURL := "https://quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}
