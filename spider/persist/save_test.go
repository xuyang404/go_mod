package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"spider/engine"
	"spider/model"
	"testing"
)

func TestItemSaver(t *testing.T) {

	client, err := elastic.NewClient(
		elastic.SetURL("http://www.vowcloud.cn:9200"),
		elastic.SetSniff(false),
	)

	profile := model.Profile{
		Name:       "你在我心中是最美",
		Birthplace: "阿克苏",
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

	const index = "dating_test"
	err = Save(client, index, item)
	if err != nil {
		panic(err)
	}

	info, err := client.Get().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual engine.Item

	err = json.Unmarshal(*info.Source, &actual)
	if err != nil {
		panic(err)
	}

	pro,err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = pro
	if actual != item {
		t.Errorf("got %+v, item %+v", actual, item)
	}

	fmt.Println(actual.Payload)
}
