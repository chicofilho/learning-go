package learninggo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing isPalindrome
func TestIsPalindrome(t *testing.T) {
	words := map[string]bool{
		"":          false,
		"a":         true,
		"aa":        true,
		"aba":       true,
		"abba":      true,
		"abcba":     true,
		"abccba":    true,
		"abcdcba":   true,
		"abcddcba":  true,
		"ra":        false,
		"raa":       false,
		"raba":      false,
		"rabba":     false,
		"rabcba":    false,
		"rabccba":   false,
		"rabcdcba":  false,
		"rabcddcba": false,
		"abc":       false,
		"aab":       false,
		"baa":       false,
		"aaa":       true,
	}
	for word, result := range words {
		assert.Equal(t, isPalindrome(word), result)
	}
}

// testing the validation. The number is passed as squared as a string
func TestValidate(t *testing.T) {
	numbers := map[int]bool{
		1:   false,
		317: true,
	}
	for num, result := range numbers {
		assert.Equal(t, validate(strconv.Itoa(num*num)), result)
	}
}

// finding the smallest palindrome
func TestFindPalindrome(t *testing.T) {
	res, ok := findSmallerPalindrome()
	assert.True(t, ok)
	assert.Equal(t, res, 836)
}
