package main

import (
	"fmt"
	"log"
	"middlewares/go_middlewares"
)


//log中间键
func Logger(request *go_middlewares.Request) {
	log.Println("请求开始")
	request.Next()
	log.Println("请求结束")
}

//捕获异常中间件
func Recover(request *go_middlewares.Request)  {
	defer func() {
		recover()
		fmt.Println("我确保panic被捕获")
	}()

	request.Next()
}

func main()  {
	request := go_middlewares.NewRequest()
	request.Register(Logger, Recover, func(request *go_middlewares.Request) {
		fmt.Println("我是业务逻辑")
	})

	request.Next()
}
