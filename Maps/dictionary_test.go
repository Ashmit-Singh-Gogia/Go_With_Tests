package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Searching in dictionary", func(t *testing.T) {
		dictionary := map[string]string{"testKey": "This the test value"}
		got := Search(dictionary, "testKey")
		want := "This the test value"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
