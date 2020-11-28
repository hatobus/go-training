package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type position struct {
	i, j int
}

type values struct {
	ax, ay, bx, by, cx, cy, dx, dy float64
}

func main() {
	f, err := os.Open("/dev/null")
	if err != nil {
		log.Fatal(err)
	}
	concurrent(8, f)
}

func concurrent(workers int, buf io.Writer) {
	if buf == nil {
		buf = os.Stdout
	}
	poschan := make(chan *position, workers)
	calcchan := make(chan *values, workers)

	for i := 0; i < workers; i++ {
		go func() {
			for p := range poschan {
				ax, ay := corner(p.i+1, p.j)
				bx, by := corner(p.i, p.j)
				cx, cy := corner(p.i, p.j+1)
				dx, dy := corner(p.i+1, p.j+1)
				if validateNumericValue(ax, ay, bx, by, cx, cy, dx, dy) {
					calcchan <- &values{ax, ay, bx, by, cx, cy, dx, dy}
				} else {
					calcchan <- nil
				}
			}
		}()
	}

	go func() {
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				poschan <- &position{i, j}
			}
		}
		close(poschan)
	}()

	fmt.Fprintf(buf, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells*cells; i++ {
		points := <-calcchan
		if points != nil {
			fmt.Fprintf(buf, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				points.ax, points.ay, points.bx, points.by, points.cx, points.cy, points.dx, points.dy)
		}
	}

	fmt.Fprintln(buf, "</svg>")
}

func naive(buf io.Writer) {
	if buf == nil {
		buf = os.Stdout
	}

	fmt.Fprintf(buf, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if validateNumericValue(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Fprintf(buf, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Fprintln(buf, "</svg>")
}

func validateNumericValue(values ...float64) bool {
	for _, v := range values {
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return false
		}
	}
	return true
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
