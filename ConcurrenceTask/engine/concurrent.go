package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func ( e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i ++  {
		createWorkder(out, e.Scheduler)
	}

	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item :=  range result.Items {
			//fmt.Printf("Got item: %v", item)
			log.Printf("Got Item %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorkder(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for  {
			s.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}