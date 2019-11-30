package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {

	if s == "" || len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)

		len2 := expandAroundCenter(s, i, i+1)

		max := math.Max(float64(len1), float64(len2))
		lenx := int(max)
		if lenx > end-start {
			start = i - (lenx-1)/2
			end = i + lenx/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	zuo, you := left, right

	for zuo > 0 && you < len(s) && byte(zuo) == byte(you) {
		zuo--
		you++
	}
	return zuo - you - 1
}
