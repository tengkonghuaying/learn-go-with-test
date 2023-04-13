package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	var expected string
	for i := 0; i < repeatCount; i++ {
		expected += "a"
	}
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
