package scheduler

import "GoSpider/ConcurrenceTask/engine"

type QueuedScheduler struct {
	 requestChan chan engine.Request
	 WorkChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request){

}
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.WorkChan <- w
}

func (s *QueuedScheduler) Run() {
	s.WorkChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workQ    []chan engine.Request

		for  {
			var activeRequest engine.Request
			var activeWorkder chan engine.Request
			if len(requestQ) >0 && len(workQ) >0  {
				activeWorkder = workQ[0]
				activeRequest = requestQ[0]
			}

			select {
				case r := <- s.requestChan:
					requestQ = append(requestQ, r)
				case w := <- s.WorkChan:
					workQ = append(workQ,w)
				case activeWorkder <-  activeRequest:
					workQ = workQ[1:]
					requestQ = requestQ[1:]
			}
		}
	}()
}

