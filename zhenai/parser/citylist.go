package parser

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
	"regexp"
)

var cityListReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(content []byte) engine.ParseResult {
	matches := cityListReg.FindAllStringSubmatch(string(content), -1)
	var result = engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       m[1],
			ParseFunc: ParseCity,
		})
		result.Items = append(result.Items, m[2])
	}
	return result
}
