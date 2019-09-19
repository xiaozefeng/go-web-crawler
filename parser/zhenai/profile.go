package zhenai

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/model/zhenai"
	"log"
)

func ParseProfile(content []byte) engine.ParseResult {
	var r engine.ParseResult
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Printf("goquery new document error:%v", err)
	}

	s1 := make([]string, 0)
	s2 := make([]string, 0)
	document.Find(".purple-btns div").Each(func(i int, s *goquery.Selection) {
		s1 = append(s1, s.Text())
	})
	document.Find(".pink-btns div").Each(func(i int, s *goquery.Selection) {
		s2 = append(s2, s.Text())
	})
	profile := zhenai.Profile{
		Name:       "",
		Gender:     "",
		Age:        s1[1],
		Height:     s1[3],
		Weight:     s1[4],
		Income:     s1[6],
		Marriage:   s1[0],
		Education:  s1[8],
		Occupation: s1[7],
		Hukou:      s2[1],
		House:      s2[5],
		Car:        s2[6],
	}

	r.Items = append(r.Items, profile)
	return r
}
