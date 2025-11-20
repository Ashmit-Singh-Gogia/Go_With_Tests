package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Searching in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"testKey": "This the test value"}

		got, _ := dictionary.Search("testKey")
		want := "This the test value"
		assertStrings(t, got, want)
	})
	t.Run("Searching a non existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("unknownKey")

		if err == nil {
			t.Fatal("Expected to get an error")
		}
		assertStrings(t, err.Error(), "could not find the word you were looking for")
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
