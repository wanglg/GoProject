package scheduler

import "crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}
func (s *QueueScheduler) WorkReady(c chan engine.Request) {
	s.workerChan <- c
}
func (s *QueueScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
}
func (s *QueueScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ [] chan engine.Request
		for {
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)

			}
		}
	}()
}
