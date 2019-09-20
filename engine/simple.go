package engine

import (
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		var curRequest = requests[0]
		requests = requests [1:]

		parseResult, err := worker(curRequest)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item:%#v", item)
		}
	}
}

func  worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url:%s", r.Url)
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch url: %s, error: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(content), nil
}
