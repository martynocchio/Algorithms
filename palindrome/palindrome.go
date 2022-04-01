package main

import (
	"strconv"
)

func CheckPalindrome(str string) bool {
	reverseString := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseString += string(str[i])
	}

	if str == reverseString {
		return true
	}
	return false
}

func CheckIntPalindrome(i int) bool {
	iToString := strconv.Itoa(i)

	reverseString := ""
	for i := len(iToString) - 1; i >= 0; i-- {
		reverseString += string(iToString[i])
	}
	if reverseString == iToString {
		return true
	}
	return false
}
