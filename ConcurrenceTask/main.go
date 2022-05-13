package main

import (
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/scheduler"
	"GoSpider/ConcurrenceTask/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Type : "url",
		Url: "http://www.zhenai.com/zhenghun",
		ParserFuc: parser.ParseCityList,
	})
}




