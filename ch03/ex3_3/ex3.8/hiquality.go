package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin = -2, -2
	xmax, ymax = 2, 2

	width, height = 1024, 1024

	iterations = 100
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for ypixel := 0; ypixel < height; ypixel++ {
		y := float64(ypixel)/height*(ymax-ymin) + ymin
		for xpixel := 0; xpixel < width; xpixel++ {
			x := float64(xpixel)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(xpixel, ypixel, MandelbrotRat(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func MandelbrotComplex64(z complex128) color.Color {
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			if n%2 == 1 {
				return color.RGBA{0xff, 0x00, 0x00, 0xff}
			} else {
				return color.RGBA{0x00, 0xff, 0x00, 0xff}
			}
		}
	}
	return color.Black
}

func MandelbrotComplex128(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n%2 == 1 {
				return color.RGBA{0xff, 0x00, 0x00, 0xff}
			} else {
				return color.RGBA{0x00, 0xff, 0x00, 0xff}
			}
		}
	}
	return color.Black
}

func MandelbrotBigFloat(z complex128) color.Color {
	realZ := (&big.Float{}).SetFloat64(real(z))
	imageZ := (&big.Float{}).SetFloat64(imag(z))

	var r, i = &big.Float{}, &big.Float{}

	// 複素数 (r+i)^2 は r^2 + 2ri + i^2 と展開できる
	for n := uint8(0); n < iterations; n++ {

		tr, ti := &big.Float{}, &big.Float{}
		tr.Mul(r, r).Sub(tr, (&big.Float{}).Mul(i, i)).Add(tr, realZ)
		ti.Mul(r, i).Mul(ti, big.NewFloat(2)).Add(ti, imageZ)
		r, i = tr, ti
		sum := &big.Float{}
		sum.Mul(r, r).Add(sum, (&big.Float{}).Mul(i, i))

		if sum.Cmp(big.NewFloat(4)) == 1 {
			if n%2 == 1 {
				return color.RGBA{0xff, 0x00, 0x00, 0xff}
			} else {
				return color.RGBA{0x00, 0xff, 0x00, 0xff}
			}
		}
	}
	return color.Black
}

func MandelbrotRat(z complex128) color.Color {
	realZ := (&big.Rat{}).SetFloat64(real(z))
	imageZ := (&big.Rat{}).SetFloat64(imag(z))

	var r, i = &big.Rat{}, &big.Rat{}

	// 複素数 (r+i)^2 は r^2 + 2ri + i^2 と展開できる
	for n := uint8(0); n < iterations; n++ {

		tr, ti := &big.Rat{}, &big.Rat{}
		tr.Mul(r, r).Sub(tr, (&big.Rat{}).Mul(i, i)).Add(tr, realZ)
		ti.Mul(r, i).Mul(ti, big.NewRat(2, 1)).Add(ti, imageZ)
		r, i = tr, ti
		sum := &big.Rat{}
		sum.Mul(r, r).Add(sum, (&big.Rat{}).Mul(i, i))

		if sum.Cmp(big.NewRat(4, 1)) == 1 {
			if n%2 == 1 {
				return color.RGBA{0xff, 0x00, 0x00, 0xff}
			} else {
				return color.RGBA{0x00, 0xff, 0x00, 0xff}
			}
		}
	}
	return color.Black
}
