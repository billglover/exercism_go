package armstrong

import (
	"math/rand"
	"testing"
)

func TestArmstrong(t *testing.T) {
	for _, tc := range testCases {
		if actual := IsNumber(tc.input); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkArmstrong10(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	ints := r.Perm(10)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := range ints {
			IsNumber(ints[i])
		}
	}
}

func BenchmarkArmstrong100(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	ints := r.Perm(100)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := range ints {
			IsNumber(ints[i])
		}
	}
}

func BenchmarkArmstrong1000(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	ints := r.Perm(1000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := range ints {
			IsNumber(ints[i])
		}
	}
}
