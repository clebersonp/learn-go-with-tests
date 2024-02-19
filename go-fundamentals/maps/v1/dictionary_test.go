package v1

import "testing"

func TestSearch(t *testing.T) {
	// map syntax -> map[key-type]value-type{initialization-key:value}
	dictionary := map[string]string{"test": "this is just a test"}
	key := "test"
	got := Search(dictionary, key)
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given key %q", got, want, key)
	}
}
