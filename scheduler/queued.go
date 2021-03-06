package scheduler

import (
	"github.com/xiaozefeng/go-web-crawler/engine"
)

type Worker chan engine.Request

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan Worker
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan Worker)
	go func() {
		var requestQ []engine.Request
		var workerQ []Worker
		for {
			var activeRequest engine.Request
			var activeWorker Worker
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}
	}()
}
