package engine

import (
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		var curRequest = requests[0]
		requests = requests [1:]

		log.Printf("Fetching url:%s", curRequest.Url)
		content, err := fetcher.Fetch(curRequest.Url)
		if err != nil {
			log.Printf("fetch url: %s, error: %v", curRequest.Url, err)
			continue
		}
		parseResult := curRequest.ParseFunc(content)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item:%#v", item)
		}
	}
}
