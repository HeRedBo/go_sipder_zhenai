package main

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/crawler_distributed/persist"
	"GoSpider/ConcurrenceTask/crawler_distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	// 打印异常并强制退出
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverServicePort),config.ElasticIndex))
}


func serveRpc(host string, index string) error {
	client ,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,&persist.ItemSaveService{
		Client: client,
		Index: index,
	})
}
