package palindromeChecker

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func IsAPalindrome(st string) bool {
	if st == "" {
		return false
	}
	st = strings.ToLower(st)
	stR := reverse(st)
	fmt.Println("Checking if", st, "is a palindrome")
	return st == stR
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
