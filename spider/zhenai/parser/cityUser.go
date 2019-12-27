package parser

import (
	"regexp"
	"spider/engine"
)



var cityUserRe = regexp.MustCompile(
	`<tr><th><a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a></th></tr>.*?<tr><td .*?><span class="grayL">性别：</span>(.*?)</td> .*?</tr>`)

var cityUrlRe = regexp.MustCompile(`<a .*? href="(http://www.zhenai.com/zhenghun/[^"]+)">([^<]+)</a>`)

func cityUserParse(contents []byte, url string, cityName string) engine.ParseResult {
	all := cityUserRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _,a := range all{
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(a[1]),
			Parser: NewProfileParse(string(a[2]), cityName, string(a[3])),
		})
	}

	urls := cityUrlRe.FindAllSubmatch(contents, -1)

	for _,url := range urls{
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(url[1]),
			Parser: NewCityUserParse(string(url[2])),
		})
	}

	return result
}