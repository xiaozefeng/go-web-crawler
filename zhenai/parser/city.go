package parser

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(content [] byte) engine.ParseResult {
	matches := cityRe.FindAllStringSubmatch(string(content), -1)
	var result engine.ParseResult
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       m[1],
			ParseFunc: engine.NiParser,
		})
		result.Items = append(result.Items, m[2])
	}
	return result
}
