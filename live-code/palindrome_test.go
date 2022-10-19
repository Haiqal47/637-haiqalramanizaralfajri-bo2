package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome_Valid_Palindrome(t *testing.T) {
	result := Palindrome("kasur ini rusak")

	assert.Equal(t, "kasur ini rusak", result, "Result must be same")

	fmt.Println("TestPalindrome_Valid_Palindrome eksekusi terhenti")
}
