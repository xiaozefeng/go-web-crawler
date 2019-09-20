package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	e.Scheduler.Run()


	var out = make(chan ParseResult)

	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler, out)
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

func createWorker(s Scheduler, out chan ParseResult) {
	in := make (chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
