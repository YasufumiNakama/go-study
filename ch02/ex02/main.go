package main

import (
	"bufio"
	"ex02/convs"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	t = flag.Bool("t", false, "templature")
	l = flag.Bool("l", false, "length")
	w = flag.Bool("w", false, "weight")
)

func main() {
	flag.Parse()
	// コマンドライン引数 or 標準入力
	if flag.NFlag() > 0 {
		arg := flag.Arg(0)
		input, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "main: %v\n", err)
			os.Exit(1)
		}
		PrintConv(input, *t, *l, *w)
	} else {
		var conv string
		var inputString string
		var input float64
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("convert [t, l, w]: ")
		if scanner.Scan() {
			conv = scanner.Text()
		}
		fmt.Printf("input: ")
		if scanner.Scan() {
			inputString = scanner.Text()
		}
		input, err := strconv.ParseFloat(inputString, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "main: %v\n", inputString)
			os.Exit(1)
		}
		PrintConv(input, conv == "t", conv == "l", conv == "w")
	}
}

func PrintConv(input float64, t bool, l bool, w bool) {
	if t {
		f := convs.Fahrenheit(input)
		c := convs.Celsius(input)
		fmt.Printf("%s = %s, %s = %s\n", f, convs.FToC(f), c, convs.CToF(c))
	}
	if l {
		f := convs.Feet(input)
		m := convs.Metre(input)
		fmt.Printf("%s = %s, %s = %s\n", f, convs.FeetToMetre(f), m, convs.MetreToFeet(m))
	}
	if w {
		p := convs.Pound(input)
		k := convs.Kilogram(input)
		fmt.Printf("%s = %s, %s = %s\n", p, convs.PoundToKilogram(p), k, convs.KilogramToPound(k))
	}
}
