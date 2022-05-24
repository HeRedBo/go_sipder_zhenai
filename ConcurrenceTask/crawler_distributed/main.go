package main

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	ItemSaver "GoSpider/ConcurrenceTask/crawler_distributed/persist/client"
	worker  "GoSpider/ConcurrenceTask/crawler_distributed/worker/client"
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/scheduler"
	"GoSpider/ConcurrenceTask/zhenai/parser"
	"fmt"
)

func main() {
	// 初始化 rpc  链接
	itemChan,err := ItemSaver.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverServicePort))
	if err != nil {
		panic(err)
	}
	processor , err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Type : "url",
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),
	})
}




