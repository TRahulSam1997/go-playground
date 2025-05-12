package main

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))

}

func isPalindrome(s string) bool {
	formattedStr := nonAlphanumericRegex.ReplaceAllString(strings.ToLower(strings.ReplaceAll(s, " ", "")), "")

	end := len(formattedStr) - 1
	for start := 0; start < len(formattedStr); start++ {
		if(formattedStr[start] != formattedStr[end]) {
			return false
		}
		end--
	}

	return true
}

