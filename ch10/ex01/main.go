package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var (
	format = flag.String("format", "jpeg", "format")
)

func main() {
	flag.Parse()
	switch *format {
	case "jpeg":
		err := toJPEG(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "toJPEG failed: %v\n", err)
			os.Exit(1)
		}
	case "png":
		err := toPNG(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "toPNG ailed: %v\n", err)
			os.Exit(1)
		}
	case "gif":
		err := toGIF(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "toGIF failed: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "format should be jpeg or png or gif\n")
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, &gif.Options{NumColors: 256})
}
