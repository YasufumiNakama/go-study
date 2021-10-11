package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	img = superSampling(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func superSampling(img *image.RGBA) *image.RGBA {
	/* 個々の画素内の複数の点のカラー値を計算して平均を求めることでピクセル化の影響を薄める */
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	out := image.NewRGBA(image.Rect(0, 0, width, height))
	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			out.Set(px, py, subPixel(img, px, py))
		}
	}
	return out
}

func subPixel(img *image.RGBA, x, y int) color.RGBA {
	/* 4点のカラー値の平均を求める */
	rgbas := []color.RGBA{
		img.RGBAAt(x, y),
		img.RGBAAt(x+1, y),
		img.RGBAAt(x, y+1),
		img.RGBAAt(x+1, y+1)}
	num := float64(len(rgbas))
	var r, g, b, a float64
	for _, rgba := range rgbas {
		r += float64(rgba.R)
		g += float64(rgba.G)
		b += float64(rgba.B)
		a += float64(rgba.A)
	}
	r = r / num
	g = g / num
	b = b / num
	a = a / num
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
