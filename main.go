package main

import (
	"fmt"
	"math"
)

func longestCommonSubsequence(x string, y string) int {
	lcs := make([][]int, len(x)+1)
	for i := 0; i < len(lcs); i++ {
		lcs[i] = make([]int, len(y)+1)
	}

	for i := 1; i < len(x)+1; i++ {
		for j := 1; j < len(y)+1; j++ {
			if x[i-1] == y[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = int(math.Max(float64(lcs[i][j-1]), float64(lcs[i-1][j])))
			}
		}
	}

	return lcs[len(x)][len(y)]
}

func main() {
	x := "abbaba"
	y := "bababba"
	res := longestCommonSubsequence(x, y)
	fmt.Printf("Result: %d\n", res)
}
