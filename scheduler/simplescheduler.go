package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChan <- r
	}()
}
func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workChan = c
}
