package client

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/crawler_distributed/rpcsupport"
	"fmt"
	"log"
)

func ItemSaver(host string) (chan interface{},error) {
	out := make(chan interface{})
	client, err := rpcsupport.NewCient(host)
	if err !=nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <- out
			// call rpc to save item
			result := ""
			err = client.Call(config.ItemSaverRpc,item, &result)
			if err != nil {
				log.Printf("item save :error saving item %v : %v\n", item, err )
			}
			fmt.Printf("es save client count:%d ,item : %v\n", itemCount, item)
			itemCount ++
		}
	}()
	return out,nil
}
