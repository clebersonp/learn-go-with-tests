package v2

import "testing"

func TestSearch(t *testing.T) {

	key := "tests"
	dictionary := Dictionary{key: "this is just a test"}
	got := dictionary.Search(key)
	want := "this is just a test"
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
