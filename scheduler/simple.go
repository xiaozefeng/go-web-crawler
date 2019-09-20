package scheduler

import "github.com/xiaozefeng/go-web-crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workChan = make (chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		s.workChan <- request
	}()
}
