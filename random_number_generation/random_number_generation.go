package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGeneration(num int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ch := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			ch <- r.Intn(num)
		}
		close(ch)
	}()
	return ch
}

func main() {
	for v := range randomNumberGeneration(10) {
		fmt.Println(v)
	}
}
