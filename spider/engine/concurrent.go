package engine

import (
	"fmt"
	"log"
	"sync"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemSaver chan Item
	RequestProcess Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	ReadyWorkChan(chan Request)
}

var loc sync.RWMutex

func (c *ConcurrentEngine)Run(seeds ...Request) {
	c.Scheduler.Run()

	//读数据到外面的chan
	out := make(chan ParseResult)

	//创建多个worker协程
	for i:=0; i<c.WorkerCount; i++ {
		go c.createWorker(c.Scheduler.WorkerChan() , out, c.Scheduler)
	}

	for _,r := range seeds {
		c.Scheduler.Submit(r)
	}

	itemCount := 1
	//循环读out出来的请求
	for {
		ParseResult := <- out
		//打印items
		for _,item := range ParseResult.Items {
			go func(item Item) {
				c.ItemSaver <- item
			}(item)
		}

		//重新再塞入请求
		for _,request := range ParseResult.Requests{
			log.Printf("Fetcher url is %s", request.Url)
			is := isDuplicate(request.Url)
			if is {
				fmt.Println("URL "+ request.Url +" is visited")
				continue
			}
			c.Scheduler.Submit(request)
		}
		itemCount++
	}
}

func (c *ConcurrentEngine)createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	for {
		s.ReadyWorkChan(in)
		//读取请求
		request := <-in
		ParseResult, err := c.RequestProcess(request)
		if err != nil {
			fmt.Printf("RequestProcess err is: %v \n", err)
			continue
		}

		//把新获取的请求传出去
		out <- ParseResult
	}
}

//去重
var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {

	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}