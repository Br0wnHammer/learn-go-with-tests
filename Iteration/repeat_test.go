package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("f", 10)
	expected := "ffffffffff"

	if repeated != expected {
		t.Errorf("want %q but got %q", expected, repeated)
	}
}