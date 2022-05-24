package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
	RequestProcessor Processor
}

type Processor  func(r Request) (ParseResult, error)

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
		//createWorkder(e.Scheduler.WorkerChan(), out, e.Scheduler)
		e.createWorkder(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	ProfileCount := 0
	for {
		result := <- out
		for _, item :=  range result.Items {

			//log.Printf("Got  Item #%d %v",ProfileCount, item)
			if item.Type == "zhengai" {
				payload := item.Payload
				log.Printf("Got Member Item #%d %v",ProfileCount, payload)
				go func(v interface{}) { e.ItemChan <- v }(payload)
				ProfileCount ++
			}
			//if  _, ok := item.(model.Member); ok {
			//	log.Printf("Got Member Item #%d %v",ProfileCount, item)
			//	go func(v interface{}) { e.ItemChan <- v }(item)
			//	ProfileCount ++
			//}
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
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func (e *ConcurrentEngine) createWorkder(in chan Request,out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for  {
			ready.WorkerReady(in)
			request := <- in
			result, err := e.RequestProcessor(request)
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
	visitedUrl[url]  = true
	return false
}