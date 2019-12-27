package parser

import (
	"log"
	"regexp"
	"spider/engine"
	"spider/model"
	"strings"
)

//预编译
var infoRegx = regexp.MustCompile(`<div class="des f-cl" .*?>(.*?)</div>`)
var pinkRegx = regexp.MustCompile(`<div *? class="m-btn pink" .*?>(.*?)</div>`)

//用户信息解析器
func profileParse(contents []byte, url string, other ProfileParseArgs) engine.ParseResult {
	name := other.UserName
	cityName := other.CityName
	gender := other.Gender

	//获取基本信息
	info := extractString(contents, infoRegx)
	urls := strings.Split(url, "/")
	id := string(urls[len(urls)-1])

	log.Printf("id is %s", id)

	str := strings.Replace(info, " ", "", -1)
	strs := strings.Split(str, "|")

	//购房
	pinks := pinkRegx.FindAllSubmatch(contents, -1)

	haveHouse := "未购房"
	haveCar := "未购车"

	for _, pink := range pinks {
		switch string(pink[1]) {
		case "已购房":
			haveHouse = "已购房"
		case "已购车":
			haveCar = "已购车"
		}
	}

	count := len(strs)
	if count < 5 {
		a := 5 - count
		for i := 0; i <= a; i++ {
			strs = append(strs, "未知")
		}

	}

	profile := model.Profile{
		Name:       name,
		City:       cityName,
		Gender:     gender,
		Birthplace: strs[0],
		Age:        strs[1],
		Education:  strs[2],
		Marriage:   strs[3],
		Height:     strs[4],
		Income:     strs[5],
		HaveHouse:  haveHouse,
		HaveCar:    haveCar,
		Url:        url,
	}

	parseResult := engine.ParseResult{
		Items: []engine.Item{{
			Url:     url,
			Id:      id,
			Type: "zhenai",
			Payload: profile,
		},},
	}

	return parseResult
}

func extractString(contents []byte, re *regexp.Regexp) string {
	//ioutil.WriteFile("test2.html", contents, 0777)

	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
