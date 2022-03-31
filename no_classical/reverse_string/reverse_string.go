package main

import "fmt"

func reverseString(str string) string {
	r := []rune(str)
	var res []rune

	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}

func main() {
	fmt.Println(reverseString("hello world!"))
}
