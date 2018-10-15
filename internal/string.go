package internal

import "strings"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReplaceRight(s string, old string, new string, n int) string {
	newRev := Reverse(new)
	oldRev := Reverse(old)
	rev := Reverse(s)
	rev = strings.Replace(rev, oldRev, newRev, n)
	return Reverse(rev)
}
