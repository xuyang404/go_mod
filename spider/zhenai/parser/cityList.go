package parser

import (
	"regexp"
	"spider/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/.*?)" .*?>(.*?)</a>`
func CityListParse(contents []byte, url string) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	all := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _,a := range all{
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(a[1]),
			Parser: NewCityUserParse(string(a[2])) ,
		})
	}

	return result
}
