package main

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/parser/zhenai"
	"github.com/xiaozefeng/go-web-crawler/scheduler"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: zhenai.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount:10,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCityList,
	})

	//engine.Run(engine.Request{
	//	Url:       "http://www.3o2o.com/sirenjing",
	//	ParseFunc: deadman.ParseTableOfContent,
	//})
}
