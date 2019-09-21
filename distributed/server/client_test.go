package main

import (
	"github.com/xiaozefeng/go-web-crawler/distributed/rpcsupport"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/model/zhenai"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "testing")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	var item = engine.Item{
		Id:   "1855225565",
		Url:  "https://album.zhenai.com/u/1855225564",
		Type: "zhenai",
		Payload: zhenai.Profile{
			Name:          "风雨相随",
			Gender:        "男士",
			Age:           "37岁",
			Height:        "173cm",
			Weight:        "75kg",
			Income:        "5-8千",
			Marriage:      "离异",
			Education:     "大专",
			Occupation:    "医生",
			Hukou:         "上海",
			Constellation: "摩羯座",
			House:         "和家人同住",
			Car:           "未买车",
			Avatar:        "https://photo.zastatic.com/images/photo/463807/1855225564/1430605127695030.jpg?scrop=1&crop=1&cpos=north&w=200&h=200",
		},
	}
	result:=""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s, err:%v", result, err)
	}

}
