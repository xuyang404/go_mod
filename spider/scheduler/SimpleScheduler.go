package scheduler

import (
	"spider/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ReadyWorkChan(chan engine.Request) {
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler)Submit(request engine.Request) {
	//并发request，解堵
	go func() {
		s.workerChan <- request
	}()
}