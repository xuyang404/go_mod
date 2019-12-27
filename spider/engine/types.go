package engine

type ParserFunc func (content []byte, url string) ParseResult

type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize() (name string, Args interface{})
}

//请求对象
type Request struct {
	Url string //url
	Parser  Parser//解析函数,
}

//解析结果
type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url string
	Id  string
	Type string
	Payload interface{}
}

type FuncParser struct{
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(content []byte, url string) ParseResult {
	return f.parser(content, url)
}

func (f *FuncParser) Serialize() (name string, Args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

