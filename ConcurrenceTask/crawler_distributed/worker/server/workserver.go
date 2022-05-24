package main

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/crawler_distributed/rpcsupport"
	"GoSpider/ConcurrenceTask/crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),&worker.CrawlerService{}))
}


