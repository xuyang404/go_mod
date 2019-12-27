package engine

import (
	"log"
	"spider/fetcher"
)

func Worker(r Request) (ParseResult,error)  {

	b,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error is %v, Fetching url is %s", err, r.Url)
		return ParseResult{}, err
	}

	//采用解析器进行解析并返回url列表和item
	return r.Parser.Parse(b, r.Url), err
}