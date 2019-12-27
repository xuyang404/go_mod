package main

import (
	"fmt"
	"spider/crawl_distributed/rpcsupport"
	"spider/crawl_distributed/worker"
	"spider/zhenai/parser"
	"testing"
	"time"
)

func TestWorkerClient(t *testing.T)  {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second*1)

	client,err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url:   "http://album.zhenai.com/u/1851500817",
		Parse: worker.SerializedParser{
			FunctionName: "123",
			Args:         parser.ProfileParseArgs{
				UserName: "你在我心中是最美",
				CityName: "阿克苏",
				Gender:   "男",
			},
		},
	}
	var result worker.ParseResult
	err = client.Call("CrawlService.Process", req, &result)
	if err != nil{
		t.Errorf("call err is : %v", err)
	}else{
		fmt.Printf("result is : %v", result)
	}
}
