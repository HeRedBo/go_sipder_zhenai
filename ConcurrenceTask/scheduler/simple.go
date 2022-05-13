package scheduler

import "GoSpider/ConcurrenceTask/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//s.workerChan <- r
	go func() { s.workerChan <- r }()

}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}



