package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just test"}

	t.Run("Known word", func(t *testing.T) {

		// got := Search(dictionary, "test")
		got, _ := dictionary.Search("test")
		want := "this is just test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatalf("expected to get an error")
		}

		assertError(t, err, want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
