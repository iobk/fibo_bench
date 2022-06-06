package fibo

import "testing"

func TestFibo0(t *testing.T) {
	fibo0(20)
}
func TestFibo1(t *testing.T) {
	fibo1(20)
}
func TestFibo2(t *testing.T) {
	fibo2(20)
}
func TestFibo3(t *testing.T) {
	fibo3(20)
}
func BenchmarkFibo0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo0(20)
	}
}
func BenchmarkFibo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo1(20)
	}
}
func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo2(20)
	}
}
func BenchmarkFibo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo3(20)
	}
}
