package main

import (
	"spider/crawl_distributed/rpcsupport"
	"spider/engine"
	"spider/model"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	const host = ":1234"
	//开启ServerRpc
	go ServeRpc(host,"test1")

	//等待服务器起来
	time.Sleep(time.Second)

	//创建客户端并连接
	client,err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

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

	//远程调用服务方法
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok"{
		t.Errorf("err: %v, result: %s", err, result)
	}

}
