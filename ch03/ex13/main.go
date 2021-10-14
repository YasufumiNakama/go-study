package main

import "fmt"

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("TB: %d\n", TB)
	fmt.Printf("PB: %d\n", PB)
	fmt.Printf("EB: %d\n", EB)
	// fmt.Printf("ZB: %d\n", ZB) // overflows int
	// fmt.Printf("YB: %d\n", YB) // overflows int
}