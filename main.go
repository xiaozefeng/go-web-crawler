package main

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/parser/deadman"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: zhenai.ParseCityList,
	//})

	engine.Run(engine.Request{
		Url:       "http://www.3o2o.com/sirenjing",
		ParseFunc: deadman.ParseTableOfContent,
	})
}
