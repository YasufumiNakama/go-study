package main

import (
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

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	// At first, determine zmax, zmin
	zmax := float64(-height)
	zmin := float64(height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az := corner(i+1, j)
			_, _, bz := corner(i, j)
			_, _, cz := corner(i, j+1)
			_, _, dz := corner(i+1, j+1)
			// 不正なポリゴン(NaN or Infを含む)をスキップする
			if hasNan([]float64{az, bz, cz, dz}) || hasInf([]float64{az, bz, cz, dz}) {
				continue
			}
			z := (az + bz + cz + dz) / 4 // 中心を基準に色を決める
			if zmax < z {
				zmax = z
			} else if zmin > z {
				zmin = z
			}
		}
	}
	// Create colored polygon
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			// 不正なポリゴン(NaN or Infを含む)をスキップする
			if hasNan([]float64{ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz}) || hasInf([]float64{ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz}) {
				continue
			}
			z := (az + bz + cz + dz) / 4 // 中心を基準に色を決める
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, zcolor(z, zmax, zmin))
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

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func zcolor(z, zmax, zmin float64) string {
	r := int((z - zmin) / (zmax - zmin) * 255)
	g := 0
	b := 255 - r
	return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
}
