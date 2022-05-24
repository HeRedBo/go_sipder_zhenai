package main

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/crawler_distributed/rpcsupport"
	"GoSpider/ConcurrenceTask/crawler_distributed/worker"
	"testing"
	"time"
)

func TestWorkerServer(t *testing.T) {

	const host = ":9000"

	go rpcsupport.ServeRpc(host,&worker.CrawlerService{})
	time.Sleep(1 *time.Second)

	client, err := rpcsupport.NewCient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request {
		Url: "https://www.zhenai.com/zhenghun/aba/",
		Type: "url",
		Text: "",
		Parser: worker.SerializedParser {
			Args: "",
			FuncName: config.ParseCityUserList,
		},
	}

	var result worker.ParserResult
	if err := client.Call("CrawlerService.Process", req, &result); err != nil {
		t.Error(err)
	}

	t.Logf("parseResult:%v\n", result)




}
