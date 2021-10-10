package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		width, err := strconv.Atoi(r.FormValue("width"))
		if err != nil {
			width = 600 // default value
		}
		height, err := strconv.Atoi(r.FormValue("height"))
		if err != nil {
			height = 320 // default value
		}
		red, err := strconv.Atoi(r.FormValue("red"))
		if err != nil {
			red = 0 // default value
		}
		green, err := strconv.Atoi(r.FormValue("green"))
		if err != nil {
			green = 0 // default value
		}
		blue, err := strconv.Atoi(r.FormValue("blue"))
		if err != nil {
			blue = 0 // default value
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		polygon(w, width, height, red, green, blue)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func polygon(out io.Writer, width, height, r, g, b int) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			// 不正なポリゴン(NaN or Infを含む)をスキップする
			if hasNan([]float64{ax, ay, bx, by, cx, cy, dx, dy}) || hasInf([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%d,%d,%d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func hasInf(vals []float64) bool {
	for _, val := range vals {
		if math.IsInf(val, 0) {
			return true
		}
	}
	return false
}

func hasNan(vals []float64) bool {
	for _, val := range vals {
		if math.IsNaN(val) {
			return true
		}
	}
	return false
}

func corner(i, j, width, height int) (float64, float64) {
	xyscale := float64(width / 2 / xyrange) // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
