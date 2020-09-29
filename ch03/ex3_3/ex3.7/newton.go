package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin = -2, -2
	xmax, ymax = 2, 2

	width, height = 1024, 1024
)

var colors = []color.RGBA{
	color.RGBA{0xfc, 0xba, 0x03, 0xff},
	color.RGBA{0xfc, 0x49, 0x03, 0xff},
	color.RGBA{0x03, 0x3d, 0xfc, 0xff},
	color.RGBA{0x03, 0x6f, 0xfc, 0xff},
}

var choosedColor = map[complex128]color.RGBA{}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for ypixel := 0; ypixel < height; ypixel++ {
		y := float64(ypixel)/height*(ymax-ymin) + ymin
		for xpixel := 0; xpixel < width; xpixel++ {
			x := float64(xpixel)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(xpixel, ypixel, zpow4_1_newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func zpow4_1_newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(z4(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := choosedColor[root]
			if !ok {
				c = colors[0]
				colors = colors[1:]
				choosedColor[root] = c
			}
			// Convert to YCbCr to make producing different shades easier.
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

// f(x) = x^4 - 1

// ニュートン法は
// x_{n+1} = x_{n} - f(x_{n}) / f'(x_{n})
// になるので

// z' = z - ( z - 1/z^3) / 4
func z4(z complex128) complex128 {
	return z*z*z*z - 1
}

func fPrime(z complex128) complex128 {
	return (z - 1/(z*z*z)) / 4
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
