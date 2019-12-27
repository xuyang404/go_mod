package engine

import (
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine)Run(seeds ...Request)  {
	//维护一个队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		//取出第一个进行处理
		r := requests[0]
		//剔除掉处理了的那个
		requests = requests[1:]

		parseRequest,err := Worker(r)
		if err != nil {
			continue
		}
		//将解析到的后续需要访问的url列表解包塞进队列
		requests = append(requests, parseRequest.Requests...)

		//打印items
		for _,item := range parseRequest.Items {
			log.Println("items is", item)
		}
	}
}

