package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(str string) string {

	if strings.Contains(str, "AM") {
		str = strings.Trim(str, "AM")
		if str[:2] == "12" {
			str = "00" + str[2:]
		}
	} else {
		trimSuffix, error := strconv.Atoi(str[:2])
		if error != nil {
			panic(error)
		}

		if trimSuffix < 12 {
			trimSuffix += 12
		}

		trimSuffixString := strconv.Itoa(trimSuffix)

		str = trimSuffixString + str[2:]
		str = strings.Trim(str, "PM")
	}
	return str
}

func main() {
	str := "12:40:22AM"
	fmt.Println(timeConversion(str))
}
