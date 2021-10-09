package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)                                // sin(30°), cos(30°)
var obj = flag.String("obj", "default", "select object from eggbox, bump, saddle") // 鶏卵の箱, こぶ, 鞍(くら)

func main() {
	flag.Parse()
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, *obj)
			bx, by := corner(i, j, *obj)
			cx, cy := corner(i, j+1, *obj)
			dx, dy := corner(i+1, j+1, *obj)
			// 不正なポリゴン(NaN or Infを含む)をスキップする
			if hasNan([]float64{ax, ay, bx, by, cx, cy, dx, dy}) || hasInf([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
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

func corner(i, j int, obj string) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y, obj)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64, obj string) float64 {
	if obj == "eggbox" {
		return 0.3 * (math.Pow(math.Sin(x/2), 2) + math.Pow(math.Sin(y/2), 2))
	} else if obj == "bump" {
		return 0.05 * (math.Pow(math.Sin(x/2), 2) + math.Pow(math.Sin(y), 2))
	} else if obj == "saddle" {
		return 0.001 * (2*x*x - 3*y*y)
	} else {
		r := math.Hypot(x, y) // distance from (0,0)
		return math.Sin(r) / r
	}
}
