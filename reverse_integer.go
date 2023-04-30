package main

import (
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	result := 0

	isNegative := 1
	abs := x
	if x < 0 {
		abs = x * -1
		isNegative = -1
	}

	madeString := strconv.Itoa(abs)
	prepared := []string{}
	for i := len(madeString) - 1; i >= 0; i-- {
		prepared = append(prepared, string(madeString[i]))
	}
	result, _ = strconv.Atoi(strings.Join(prepared, ""))
	result = result * isNegative
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
