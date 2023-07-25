package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithZeroIteration(t *testing.T) {
	repeated := Repeat("a", 0)
	expected := "a"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithNegativeIteration(t *testing.T) {
	repeated := Repeat("a", -10)
	expected := "The function can't support negative iteration"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Writing Benchmarks -> another first-class feature of the language and its similar to writing tests.

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
