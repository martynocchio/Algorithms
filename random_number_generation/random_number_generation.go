package main

import (
	"fmt"
	"math/rand"
)

func randNumsGenerator(n int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			r := rand.Intn(n)
			out <- r
		}
		close(out)
	}()
	return out
}

func main() {
	for num := range randNumsGenerator(10) {
		fmt.Println(num)
	}
}
