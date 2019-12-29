package main

//生产者，负责向管道投递数据
func producer(n int) <-chan int {
	out := make(chan int)
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
	out := make(chan int)

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

func main() {
	in := producer(10000000)
	ch := square(in)

	//消费者，从管道里消费处理好的数据
	for _ = range ch {
		//fmt.Println(c)
	}
}
