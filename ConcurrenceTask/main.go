package main

import (
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/persist"
	"GoSpider/ConcurrenceTask/scheduler"
	"GoSpider/ConcurrenceTask/zhenai/parser"
)

func main() {
	// 初始化 es 链接

	itemChan,err := persist.ItemSaver("dating_prifile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
	}
	e.Run(engine.Request{
		Type : "url",
		Url: "http://www.zhenai.com/zhenghun",
		ParserFuc: parser.ParseCityList,
	})
}




