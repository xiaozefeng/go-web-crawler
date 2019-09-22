package main

import (
	"github.com/xiaozefeng/go-web-crawler/distributed/config"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/parser/zhenai"
	"github.com/xiaozefeng/go-web-crawler/persist"
	"github.com/xiaozefeng/go-web-crawler/scheduler"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: zhenai.ParseCityList,
	//})

	itemChan, err := persist.ItemSaver(config.ElasticSearchIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 10,
		ItemChan:  itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/shanghai",
	//	ParseFunc: zhenai.ParseCityWithGoQuery,
	//})

	//engine.Run(engine.Request{
	//	Url:       "http://www.3o2o.com/sirenjing",
	//	ParseFunc: deadman.ParseTableOfContent,
	//})
}
