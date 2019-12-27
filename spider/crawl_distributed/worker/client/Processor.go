package client

import (
	"net/rpc"
	"spider/crawl_distributed/config"
	"spider/crawl_distributed/worker"
	"spider/engine"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor)  {
	return func(req engine.Request) (engine.ParseResult, error) {
		var result  worker.ParseResult
		client := <- clientChan
		err := client.Call(config.CrawlServiceRpc, worker.SerializeRequest(req), &result)

		if err != nil {
			return engine.ParseResult{}, err
		}

		eResult := worker.DeserializeResult(result)
		return eResult,nil
	}
}
