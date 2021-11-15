package main

import (
	"fmt"
)

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	minVal := vals[0]
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal, nil
}

// 少なくとも1つの引数が必要なversion
func minAtLeastOneArg(one int, vals ...int) int {
	minVal := one
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
