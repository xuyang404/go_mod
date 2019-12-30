package main

import (
	"fmt"
	"sync"
)

//生产者，负责向管道投递数据
func producer(n int) <-chan int {
	//带缓冲channel，增加生产并发量
	out := make(chan int, 100)
	a := n / 10
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		//开一个goroutine，持续向管道投递数据
		go func() {
			for i := 0; i < a; i++ {
				out <- i
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	//返回管道
	return out
}

//操作者，负责从管道读取数据并进行逻辑操作再放回管道
func square(in <-chan int) <-chan int {
	//带缓冲channel，增加处理并发量
	out := make(chan int, 100)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		//开一个goroutine，从管道持续读取，进行乘方操作，再把读出来的数据放入新的管道
		go func() {
			for i := range in {
				out <- i * i
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	//返回新的管道
	return out
}

func main() {
	in := producer(10000000)

	//fan-out
	c1 := square(in)

	//消费者，从管道里消费处理好的数据
	for c := range c1 {
		fmt.Println(c)
	}
}
