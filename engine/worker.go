package engine

import (
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"log"
)


func  worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url:%s", r.Url)
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch url: %s, error: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(content), nil
}
