package v1

import "testing"

func TestSearch(t *testing.T) {
	// map syntax -> map[key-type]value-type{initialization-key:value}
	// key type is special. it can only be a comparable type
	// the value type, on the other hand, can be any type
	key := "tests"
	dictionary := map[string]string{key: "this is just a test"}
	got := Search(dictionary, key)
	want := "this is just a test"
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
