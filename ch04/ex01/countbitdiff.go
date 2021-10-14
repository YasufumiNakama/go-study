package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CountBitDiff(x, y [32]uint8) int {
	var count int
	for i := 0; i < 32; i++ {
		count += int(pc[x[i]^y[i]])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%d\n", c1, c2, CountBitDiff(c1, c2))
}
