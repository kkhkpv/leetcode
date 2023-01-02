package main

import "strings"

func main() {

}

func isSubsequence(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 > l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if !strings.Contains(t, string(s[i])) {
			return false
		} else {
			indx := strings.IndexByte(t, s[i])
			t = t[indx+1:]
		}
	}

	return true
}
