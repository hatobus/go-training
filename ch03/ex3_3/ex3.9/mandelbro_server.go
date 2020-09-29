package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", MandelbrotPlotHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func MandelbrotPlotHandler(w http.ResponseWriter, r *http.Request) {
	var (
		xmin, ymin = -2, -2
		xmax, ymax = 2, 2

		width, height = 1024, 1024
	)

	strheight := r.URL.Query().Get("height")
	strwidth := r.URL.Query().Get("width")

	strxmin := r.URL.Query().Get("xmin")
	strymin := r.URL.Query().Get("ymin")

	strxmax := r.URL.Query().Get("xmax")
	strymax := r.URL.Query().Get("ymax")

	if strheight == "" {
		strheight = strconv.Itoa(height)
	}

	if strwidth == "" {
		strwidth = strconv.Itoa(width)
	}

	if strxmin == "" {
		strxmin = strconv.Itoa(xmin)
	}

	if strymin == "" {
		strymin = strconv.Itoa(ymin)
	}

	if strxmax == "" {
		strxmax = strconv.Itoa(xmax)
	}

	if strymax == "" {
		strymax = strconv.Itoa(ymax)
	}

	w.WriteHeader(http.StatusCreated)

	var err error
	height, err = strconv.Atoi(strheight)
	if err != nil || height <= 0 {
		http.Error(w, "invalid height value", http.StatusBadRequest)
		return
	}

	width, err = strconv.Atoi(strwidth)
	if err != nil || width <= 0 {
		http.Error(w, "invalid width value", http.StatusBadRequest)
		return
	}

	xmin, err = strconv.Atoi(strxmin)
	if err != nil || xmin >= 0 {
		log.Println(xmin, err)
		http.Error(w, "invalid xmin value", http.StatusBadRequest)
		return
	}

	xmax, err = strconv.Atoi(strxmax)
	if err != nil || xmax <= 0 {
		http.Error(w, "invalid xmax value", http.StatusBadRequest)
		return
	}

	ymin, err = strconv.Atoi(strymin)
	if err != nil || ymin >= 0 {
		http.Error(w, "invalid ymin value", http.StatusBadRequest)
		return
	}

	ymax, err = strconv.Atoi(strymax)
	if err != nil || ymax <= 0 {
		http.Error(w, "invalid ymax value", http.StatusBadRequest)
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*float64(ymax-ymin) + float64(ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 100

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
