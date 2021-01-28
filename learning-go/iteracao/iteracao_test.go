package iteracao

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	if result != expected {
		t.Errorf("Expected: '%s', Result: '%s'", expected, result)
	}
}

// To execute benchmark type in terminal: go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ { // run benchmark N times
		Repeat("a", 5)
	}

}
