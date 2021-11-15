package main

import (
	"fmt"
)

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	maxVal := vals[0]
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}

// 少なくとも1つの引数が必要なversion
func maxAtLeastOneArg(one int, vals ...int) int {
	maxVal := one
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
