package main

import (
	"flag"
	"log"
	"spider/crawl_distributed/rpcsupport"
	"spider/crawl_distributed/worker"
)

var host = flag.String("host", ":9000", "the host port for me to listen on")

func main()  {
	flag.Parse()
	log.Fatal(rpcsupport.ServeRpc(*host, worker.CrawlService{}))
}
