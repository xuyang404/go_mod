package main

import (
	"flag"
	"log"
	"net/rpc"
	"spider/crawl_distributed/config"
	itemSaver "spider/crawl_distributed/persist/client"
	"spider/crawl_distributed/rpcsupport"
	workerProcessor "spider/crawl_distributed/worker/client"
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
	"strings"
)

var saverHost = flag.String("SaverHost", ":1234", "ItemSaver host (Separate with a comma)")
var workerHost = flag.String("WorkerHost", ":9000", "Worker host (Separate with a comma)")

func main() {
	flag.Parse()

	savers := strings.Split(*saverHost, ",")
	saversPool := createClientPool(savers)
	//item, err := persist.ItemSaver("dating_profile")
	item:= itemSaver.ItemSaver(saversPool, config.ElasticIndex, config.ItemSaverRpc)

	hosts := strings.Split(*workerHost, ",")
	pool := createClientPool(hosts)
	processor := workerProcessor.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
		ItemSaver:   item,
		RequestProcess: processor,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/",
		Parser: engine.NewFuncParser(parser.CityListParse, "CityListParser"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client  {
	var clients []*rpc.Client
	for _,host := range hosts{
		client,err := rpcsupport.NewClient(host)
		if err == nil{
			clients = append(clients, client)
			log.Printf("connect success %s\n", host)
		}else{
			log.Printf("connect fail %s: %v\n", host, err)
		}
	}

	out := make(chan *rpc.Client)

	go func() {
		for  {
			for _,c := range clients{
				out <- c
			}
		}

	}()
	return out
}
