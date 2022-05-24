package client

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/crawler_distributed/rpcsupport"
	"GoSpider/ConcurrenceTask/crawler_distributed/worker"
	"GoSpider/ConcurrenceTask/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor,error) {
	client, err := rpcsupport.NewCient(fmt.Sprintf(":%d",config.WorkerPort0))
	if err != nil {
		return nil , err
	}

	return  func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRuquest(req)
		var sResult worker.ParserResult
		err := client.Call(config.CrawlerServiceRpc,sReq,&sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult)
	},nil
}