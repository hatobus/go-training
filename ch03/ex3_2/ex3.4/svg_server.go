package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
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

func main() {

	// height, width, color を指定できる
	// colorは配列で渡され、山が1番目、谷が2番目の要素で描かれる
	// color="yellow,purple" の場合には山が黄色、谷が紫
	http.HandleFunc("/", generateSVGHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func generateSVGHandler(w http.ResponseWriter, r *http.Request) {
	strheight := r.URL.Query().Get("height")
	strwidth := r.URL.Query().Get("width")

	var colors []string
	color := r.URL.Query().Get("color")
	if color != "" {
		colors = strings.Split(color, ",")
	} else {
		colors = []string{"red", "blue"}
	}

	if strheight == "" {
		strheight = strconv.Itoa(height)
	}

	if strwidth == "" {
		strwidth = strconv.Itoa(width)
	}

	height, err := strconv.Atoi(strheight)
	if err != nil || height <= 0 {
		http.Error(w, "invalid height value", http.StatusBadRequest)
		return
	}

	width, err := strconv.Atoi(strwidth)
	if err != nil || width <= 0 {
		http.Error(w, "invalid width value", http.StatusBadRequest)
		return
	}

	if len(colors) != 2 {
		http.Error(w, "invalid color, length must be 2", http.StatusBadRequest)
		return
	}

	err = genSVG(w, height, width, colors)
	if err != nil {
		log.Println(err)
		http.Error(w, "svg generate failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusCreated)
	return
}

func genSVG(w io.Writer,height, width int, colors []string) error {
	_, err := fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	if err != nil {
		return err
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			color := resolveColor(i, j, colors[0], colors[1])

			if validateNumericValue(ax, ay, bx, by, cx, cy, dx, dy) {
				_, err := fmt.Fprintf(w, "<polygon fill='%v' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color, ax, ay, bx, by, cx, cy, dx, dy)
				if err != nil {
					return err
				}
			}
		}
	}

	if _, err = fmt.Fprintln(w, "</svg>"); err != nil {
		return err
	}

	return nil
}

func resolveColor(i, j int, colorUpper, colorLower string) string {
	min, max := math.NaN(), math.NaN()

	for xoffset := 0; xoffset <= 1; xoffset++ {
		for yoffset := 0; yoffset <= 1; yoffset++ {
			x := xyrange * (float64(i+xoffset)/cells - 0.5)
			y := xyrange * (float64(j+yoffset)/cells - 0.5)
			z := f(x, y)

			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	if math.Abs(max) > math.Abs(min) {
		return colorUpper
	} else {
		return colorLower
	}
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
