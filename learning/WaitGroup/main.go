package main

import (
	"fmt"
	"strconv"
	"time"
	"wg/wg"
)

func main() {
	wg1 := wg.NewWg(2)
	res := make(chan interface{}, 10)

	ch := wg1.Square(res)

	for i := 0; i < 10; i++ {
		wg1.Add()
		go func(i int) {
			m := make(map[string]interface{})
			m[strconv.Itoa(i)] = i
			res <- m
			time.Sleep(time.Second)
			wg1.Done()
		}(i, )
	}

	 wg1.Wait()
	close(res)

	fmt.Println("完了 ")
	fmt.Printf("%v", <-ch)

}
