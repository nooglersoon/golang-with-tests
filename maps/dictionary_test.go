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

func TestAdd(t *testing.T) {

	t.Run("New key", func(t *testing.T) {
		var key string = "key"
		var value string = "Value"
		dictionary := Dictionary{}

		err := dictionary.Add(key, value)

		assertError(t, err, nil)

		assertDefinitions(t, dictionary, key, value)

	})

	t.Run("Add into existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinitions(t, dictionary, key, value)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinitions(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	want := ErrNotFound

	if err == nil {
		t.Fatalf("expected to get an error")
	}

	assertError(t, err, want)

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

func assertDefinitions(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)

}
