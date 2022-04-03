package main

import (
	"fmt"
	"sync"
)

func joinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))
		for _, v := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for value := range ch {
					mergedCh <- value
				}
			}(v, wg)
		}
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for _, v := range []int{1, 2, 3} {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range []int{4, 5, 6} {
			ch2 <- v
		}
		close(ch2)
	}()

	go func() {
		for _, v := range []int{7, 8, 9} {
			ch3 <- v
		}
		close(ch3)
	}()

	result := 0
	for num := range joinChannels(ch1, ch2, ch3) {
		result += num
	}
	fmt.Println(result)
}

//mergedCh := make(chan int)
//
//go func() {
//	wg := &sync.WaitGroup{}
//	wg.Add(len(chs))
//
//	for _, ch := range chs {
//		go func(ch <-chan int, wg *sync.WaitGroup) {
//			defer wg.Done()
//			for v := range ch {
//				mergedCh <- v
//			}
//		}(ch, wg)
//	}
//	wg.Wait()
//	close(mergedCh)
//}()
//return mergedCh
