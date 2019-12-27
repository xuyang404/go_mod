package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	result := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			result <- i
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(result)

	fmt.Println("完了 ")
	a := make([]int, 1)
	for val := range result{
		a = append(a, val)
	}
	fmt.Printf("%#v", a)

}
