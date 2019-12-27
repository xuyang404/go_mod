package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"spider/crawl_distributed/config"
	"spider/engine"
	"spider/zhenai/parser"
)

type SerializedParser struct {
	FunctionName string
	Args         interface{}
}

type Request struct {
	Url   string
	Parse SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

//序列化请求
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parse: SerializedParser{
			FunctionName: name,
			Args:         args,
		},
	}
}

//序列化结果
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

//反序列化请求
func DeserializeRequest(r Request) (engine.Request, error) {

	p, err := deserializeParser(r.Parse)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

//反序列化Request里的Parser
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FunctionName {
	case config.CityUserParse:
		if cityName, ok := p.Args.(string); ok {
			return parser.NewCityUserParse(cityName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case config.CityListParer:
		return engine.NewFuncParser(parser.CityListParse, config.CityListParer), nil
	case config.ProfileParse:
		var args parser.ProfileParseArgs
		err := FromJsonObj(p.Args, &args)
		if err != nil {
			return nil, err
		}
		return parser.NewProfileParse(args.UserName, args.CityName, args.Gender), nil
	default:
		return nil, fmt.Errorf("unknown parser name")
	}
}

func FromJsonObj(o interface{}, d interface{}) error {
	b, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("json Marshal err: %v", err)
	}

	err = json.Unmarshal(b, d)
	if err != nil {
		return fmt.Errorf("json Unmarshal err: %v", err)
	}

	return nil
}

//反序列化结果
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("DeserializeRequest error: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}

	return result
}
