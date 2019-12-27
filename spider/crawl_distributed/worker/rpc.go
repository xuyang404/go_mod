package worker

import (
	"spider/engine"
)

type CrawlService struct{}

func (c CrawlService) Process(req Request,
	result *ParseResult) error {
	request, err := DeserializeRequest(req)
	if err != nil{
		return err
	}

	parseResult,err := engine.Worker(request)

	if err != nil {
		return err
	}

	*result = SerializeResult(parseResult)

	return nil
}
