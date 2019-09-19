package deadman

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
	"regexp"
)

var tableOfContentRe = regexp.MustCompile(`<a class="thumbnail" href="(http://www.3o2o.com/sirenjing/[a-z0-9]+.html)"[^>]*>([^<]+)</a>`)

func ParseTableOfContent(content [] byte) engine.ParseResult {
	var r engine.ParseResult
	matches := tableOfContentRe.FindAllStringSubmatch(string(content), -1)
	for _, m := range matches {
		//r.Requests = append(r.Requests, engine.Request{
		//	Url:       m[1],
		//	ParseFunc: engine.NiParser,
		//})
		r.Items = append(r.Items, m[2])
	}
	return r
}
