package zhenai

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/model/zhenai"
	"log"
)

func ParseProfile(content []byte, info zhenai.UserInfo) engine.ParseResult {
	var r engine.ParseResult
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Printf("goquery new document error:%v", err)
	}

	var profile = zhenai.Profile{}
	profile.Name = info.Name
	profile.Gender = info.Gender
	profile.Avatar = info.Avatar
	document.Find(".purple-btns div").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		switch i {
		case 1:
			profile.Age = text
		case 3:
			profile.Height = text
		case 4:
			profile.Weight = text
		case 6:
			profile.Income = text
		case 0:
			profile.Marriage = text
		case 8:
			profile.Education = text
		case 7:
			profile.Occupation = text
		}

	})
	document.Find(".pink-btns div").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		switch i {
		case 1:
			profile.Hukou = text
		case 5:
			profile.House = text
		case 6:
			profile.Car = text
		}
	})
	var item = engine.Item{
		Id:      info.Id,
		Url:     info.Url,
		Type:    "zhenai",
		Payload: profile,
	}
	r.Items = append(r.Items, item)
	return r
}
