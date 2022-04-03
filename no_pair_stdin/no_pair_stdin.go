package main

import (
	"bufio"
	"fmt"
	"os"
)

func NoPair() {
	sc := bufio.NewScanner(os.Stdin)
	hash := map[string]int{}

	for sc.Scan() {
		if _, value := hash[sc.Text()]; !value {
			hash[sc.Text()] = 1
			continue
		}
		hash[sc.Text()] += 1
	}

	for key, value := range hash {
		if value == 1 {
			fmt.Println(key)
		}
	}
}

func main() {
	NoPair()
}
