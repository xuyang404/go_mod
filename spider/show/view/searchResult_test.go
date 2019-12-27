package view

import (
	"os"
	"spider/engine"
	"spider/model"
	pmodel "spider/show/model"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	s := CreateSearchResultView("index.html")

	page := pmodel.SearchResultData{}
	page.Hits = 123
	page.PrevPage = 1
	page.NextPage = 3

	profile := model.Profile{
		Name:       "你在我心中是最美",
		Birthplace: "阿克苏",
		Gender:     "男",
		Age:        "21岁",
		Education:  "高中及以下",
		Marriage:   "未婚",
		Height:     "170cm",
		Income:     "5001-8000元",
		HaveHouse:  "未购房",
		HaveCar:    "未购车",
		Url:        "http://album.zhenai.com/u/1851500817",
	}

	item := engine.Item{
		Url:     "http://album.zhenai.com/u/1851500817",
		Id:      "1851500817",
		Type:    "zhenai",
		Payload: profile,
	}

	for i:=0; i<10; i++ {
		page.Items = append(page.Items, item)
	}

	page.Count = len(page.Items)
	w,err := os.Create("test.html")

	if err != nil {
		panic(err)
	}

	err = s.Render(w, page)

	if err != nil {
		panic(err)
	}

}
