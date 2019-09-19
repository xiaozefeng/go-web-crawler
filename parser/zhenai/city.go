package zhenai

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"log"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(content [] byte) engine.ParseResult {
	matches := cityRe.FindAllStringSubmatch(string(content), -1)
	var result engine.ParseResult
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: m[1],
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, m[2], "", "")
			},
		})
		result.Items = append(result.Items, m[2])
	}
	return result
}

func ParseCityWithGoQuery(content []byte) engine.ParseResult {
	var r engine.ParseResult
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Printf("goquey parse city error:%v\n", err)
		return r
	}

	document.Find(".g-list .list-item").Each(func(i int, s *goquery.Selection) {
		var c = s.Find(".content").First()
		gender := c.Find("tr:nth-child(2) td").First().Text()
		a := c.Find("a").First()
		name := a.Text()
		url, _ := a.Attr("href")
		var imageSrc, _ = s.Find(".photo a img").First().Attr("src")
		r.Requests = append(r.Requests, engine.Request{
			Url: url,
			ParseFunc: func(bs []byte) engine.ParseResult {
				return ParseProfile(bs, name, imageSrc, gender[9:])
			},
		})
		//r.Items = append(r.Items, name)
	})
	return r
}
