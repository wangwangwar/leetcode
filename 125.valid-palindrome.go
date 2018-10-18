package main

import "strings"
import "fmt"

func IsPalindrome(s string) bool {
	if "" == s {
		return true
	}

	runeArray := []rune(s)
	l := len(runeArray)
	head := 0
	tail := l - 1
	for {
		for {
			if head > l-1 || isAlphanumeric(runeArray[head]) {
				break
			}
			head++
		}

		for {
			if tail < 0 || isAlphanumeric(runeArray[tail]) {
				break
			}
			tail--
		}

		if head >= tail {
			return true
		}

		if strings.ToLower(string(runeArray[head])) != strings.ToLower(string(runeArray[tail])) {
			return false
		}

		head++
		tail--
	}
}

func isAlphanumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsPalindrome("A"))
	fmt.Println(IsPalindrome(""))
	fmt.Println(IsPalindrome(" "))
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("race a car"))
	fmt.Println(IsPalindrome("0P"))
}
