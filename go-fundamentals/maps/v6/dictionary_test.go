package v6

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	correctKey := "test"
	dictionary := Dictionary{correctKey: "this is just a test"}
	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search(correctKey)
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		// initialize an empty map
		dictionary := Dictionary{}
		key := "test"
		want := "this is just a test"

		err := dictionary.Add(key, want)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, want)
	})
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		want := "this is just a test"
		dictionary := Dictionary{key: want}

		err := dictionary.Add(key, "new test")
		assertError(t, err, ErrKeyExists)
		assertDefinition(t, dictionary, key, want)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, want string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added key:", err)
	}
	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}
