package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)


var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // black
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
}

const (
	blackIndex = 0
)

var optionalColorIndex = []uint8{1, 2, 3, 4}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous_color(os.Stdout)
}

func lissajous_color(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			next := optionalColorIndex[rand.Intn(len(optionalColorIndex))]
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
			next)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

