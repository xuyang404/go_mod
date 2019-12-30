package main

import (
	"sync"
)

//fan-in fan-out模式


//生产者，负责向管道投递数据
func producer(n int) <-chan int {
	//带缓冲channel，增加并发量
	out := make(chan int, 1000)
	//开一个goroutine，持续向管道投递数据
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()

	//返回管道
	return out
}

//操作者，负责从管道读取数据并进行逻辑操作再放回管道
func square(in <-chan int) <-chan int {
	out := make(chan int, 1000)

	//开一个goroutine，从管道持续读取，进行乘方操作，再把读出来的数据放入新的管道
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()

	//返回新的管道
	return out
}

//负责收集处理结果
//fan-in
func merge(in ...<-chan int) <-chan int {
	//带缓冲channel，增加并发量
	out := make(chan int, 1000)

	//用来保证所有collect跑完之后可以关闭out，否则会出现死锁
	wg := sync.WaitGroup{}

	collect := func(in <-chan int) {
		defer wg.Done()
		for i := range in {
			out <- i
		}
	}

	//遍历每个操作者返回的管道并进行处理
	for _, in := range in {
		wg.Add(1)
		go collect(in)
	}

	//错误，会导致程序卡在这，主程序读取不倒out，从而出现死锁
	//wg.Wait()
	//close(out)

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producer(10000000)

	//fan-out
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)


	//消费者，从管道里消费处理好的数据
	for _ = range merge(c1, c2, c3) {
		//fmt.Println(c)
	}
}
