package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChan <- r
	}()
}
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChan
}

//func (s *SimpleScheduler) Run() {
//
//}
