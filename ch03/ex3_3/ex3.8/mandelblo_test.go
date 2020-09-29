package main_test

import (
	"image/color"
	"testing"

	ex38 "github.com/hatobus/go-training/ch03/ex3_3/ex3.8"
)

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkComplex64(b *testing.B) {
	benchmarkMandelbrot(b, ex38.MandelbrotComplex64)
}

func BenchmarkComplex128(b *testing.B) {
	benchmarkMandelbrot(b, ex38.MandelbrotComplex128)
}

func BenchmarkBigFloat(b *testing.B) {
	benchmarkMandelbrot(b, ex38.MandelbrotBigFloat)
}

func BenchmarkRat(b *testing.B) {
	benchmarkMandelbrot(b, ex38.MandelbrotRat)
}
