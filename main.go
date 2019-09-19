package main

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
