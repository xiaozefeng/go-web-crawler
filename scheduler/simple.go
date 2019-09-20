package scheduler

import "github.com/xiaozefeng/go-web-crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(w chan engine.Request) {
	s.workChan = w
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		s.workChan <- request
	}()
}
