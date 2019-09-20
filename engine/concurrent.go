package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	var in = make(chan Request)
	e.Scheduler.ConfigureMasterWorkChan(in)
	var out = make(chan ParseResult)

	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		parseResult := <-out
		for _, item := range parseResult.Items {
			fmt.Printf("Got item:%#v", item)
		}

		for _, req := range parseResult.Requests {
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
