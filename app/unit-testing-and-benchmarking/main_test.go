package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestIntMin(t *testing.T) {
	ret := IntMin(2, -2)
	if ret != -2 {
		t.Errorf("IntMin(2, -2) = %d, should be -2", ret)
	}
}

func TestIntMinTableDiven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{1, 0, 0},
		{2, 1, 1},
		{0, 1, 0},
		{100, -9, -9},
		{30, 29, 29},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("test[%d, %d]", test.a, test.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(test.a, test.b)
			if ans != test.want {
				t.Errorf("ans: %d, want: %d", ans, test.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
