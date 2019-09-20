package engine

import (
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
