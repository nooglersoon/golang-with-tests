package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Aji")

	got := buffer.String()
	want := "Hello, Aji"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}
