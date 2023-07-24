package hello

import "testing"

// We can run multiple tests at once by grouping into subtests.
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := createHello("Chris", "")
		want := "Hello Chris!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := createHello("", "")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := createHello("Ajie", "spanish")
		want := "Hola Ajie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := createHello("Ajie", "french")
		want := "Bonjour Ajie!"
		assertCorrectMessage(t, got, want)
	})

}

// Refactor error message by extracting into a helper

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*

Basic Tests

func TestHello(t *testing.T) {
	// Declare variable
	// varName := value
	got := createHello("Fauzi")
	want := "Hello Fauzi!"

	if got != want {
		// We are calling the Errorf method on our t which will print out a message and fail the test.
		// The f stands for format which allows us to build a string with values inserted into the placeholder values %q.
		t.Errorf("got %q want %q", got, want)
	}
}

*/
