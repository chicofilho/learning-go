// solving The Palindrome-Number Problem
// http://icarus.cs.weber.edu/~dab/cs1410/textbook/8.Strings/progexample/pal-number.html
// This has the only purppose to test and learn Go
package learninggo

import (
	"math"
	"strconv"
)

// receives a string, returns true if it is a palindrome and false otherwise
func isPalindrome(word string) bool {
	if len(word) == 0 {
		return false
	}
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

// finding the smallest palindrome
func findSmallerPalindrome() (int, bool) {
	el := int(math.Sqrt(1))
	for {
		sqrStr := strconv.Itoa(el * el)
		if validate(sqrStr) && isPalindrome(sqrStr) {
			return el, true
		}
		el++
	}
}

// validate helper: expects a number string with 6 digits that doens't start or end with zero
func validate(sqrStr string) bool {
	return len(sqrStr) >= 6 && sqrStr[0] != '0' && sqrStr[len(sqrStr)-1] != '0'
}
