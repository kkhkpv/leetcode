package leetcode

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	sRuneSlice := []rune(s)
	if len(sRuneSlice) <= 1 {
		return len(sRuneSlice)
	}
	left, right, res := 0, 0, 0
	control := ""
	for right < len(sRuneSlice) {
		if !strings.ContainsRune(control, sRuneSlice[right]) {
			control += string(sRuneSlice[right])
			right++
			res = int(math.Max(float64(res), float64(len([]rune(control)))))
		} else {
			control = strings.Replace(control, string(sRuneSlice[left]), "", 1)
			left++
		}
	}
	return res
}
