package engine

import (
	"GoSpider/ConcurrenceTask/zhenai/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func ( e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i ++  {
		createWorkder(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	ProfileCount := 0
	for {
		result := <- out
		for _, item :=  range result.Items {
			//fmt.Printf("Got item: %v", item)
			//log.Printf("Got Item #%d %v", itemCount,item)
			if  _, ok := item.(model.CityProfile); ok {
				log.Printf("Got CityProfile Item #%d %v",ProfileCount, item)
				ProfileCount ++
			}
		}

		for _, request := range result.Requests {
			// url 去重
			if isDuplicate(request.Url) {
				log.Printf("Duplicate request :" + "%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func createWorkder(in chan Request,out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for  {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrl = make(map[string]bool)
func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}