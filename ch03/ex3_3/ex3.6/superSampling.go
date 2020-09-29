package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin = -2, -2
	xmax, ymax = 2, 2

	width, height = 1024, 1024

	X = 4 / width
	Y = 4 / height
)

func main() {
	offsetX := []float64{-X, X}
	offsetY := []float64{-Y, Y}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for ypixel := 0; ypixel < height; ypixel++ {
		y := float64(ypixel)/height*(ymax-ymin) + ymin
		for xpixel := 0; xpixel < width; xpixel++ {
			x := float64(xpixel)/width*(xmax-xmin) + xmin

			subPixel := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offsetX[i], y+offsetY[i])
					subPixel = append(subPixel, mandelbrot(z))
				}
			}
			img.Set(xpixel, ypixel, average(subPixel))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

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

func average(colors []color.Color) color.Color {
	var r, g, b, a uint16

	length := len(colors)
	for _, c := range colors {
		tr, tg, tb, ta := c.RGBA()
		r += uint16(tr / uint32(length))
		g += uint16(tg / uint32(length))
		b += uint16(tb / uint32(length))
		a += uint16(ta / uint32(length))
	}

	return color.RGBA64{r, g, b, a}
}
