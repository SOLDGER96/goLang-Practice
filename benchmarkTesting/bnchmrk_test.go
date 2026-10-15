package main
import "testing"

func TestAdd(t *testing.T) {
	result := add(4,7)
	if result != 11 {
		t.Errorf("Expected 11, got %d", result)
	}
}

// Benchmark Funcn
func BenchmarkAdd(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		add(4,7)
	}
}

// Use go text -bench . to tun the benchmark