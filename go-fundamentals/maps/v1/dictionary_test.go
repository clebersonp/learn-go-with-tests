package v1

import "testing"

func TestSearch(t *testing.T) {
	// map syntax -> map[key-type]value-type{initialization-key:value}
	// key type is special. it can only be a comparable type
	// the value type, on the other hand, can be any type
	dictionary := map[string]string{"test": "this is just a test"}
	key := "test"
	got := Search(dictionary, key)
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given key %q", got, want, key)
	}
}
