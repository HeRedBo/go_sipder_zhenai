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
}

func ( e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i ++  {
		createWorkder(in, out)
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

func createWorkder(in chan Request, out chan ParseResult) {
	go func() {
		for  {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}