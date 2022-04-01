package main

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		testName       string
		stringToCheck  string
		expectedResult bool
	}{
		{
			testName:       "first check",
			stringToCheck:  "helleh",
			expectedResult: true,
		},
		{
			testName:       "second check",
			stringToCheck:  "adaddsaa",
			expectedResult: false,
		},
	}

	for _, v := range testCases {
		t.Run(v.testName, func(t *testing.T) {
			result := CheckPalindrome(v.stringToCheck)
			if result == true && v.expectedResult == true {
				fmt.Println("test finished successfully, str is palindrome")
			} else if result == true && v.expectedResult == false || result == false && v.expectedResult == true {
				fmt.Println("test failed")
			} else if result == false && v.expectedResult == false {
				fmt.Println("test finished successfully, str is not palindrome")
			}
		})
	}
}
