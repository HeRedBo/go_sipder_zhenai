package persist

import (
	"GoSpider/ConcurrenceTask/zhenai/model"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"strconv"
)


// es 变量定义修改
var Type = "zhenai"

func ItemSaver(index string) (chan interface{},error) {
	out := make(chan interface{})
	// 判断索引是否存在并创建索引
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		// Handle error
		log.Println("es exists query fail:", err)
		return nil, err
		//return "",nil
	}
	if !exists {
		_, err := client.CreateIndex(index).
			Do(context.Background())
		if err != nil {
			log.Println("es CreateIndex  fail:", err)
			// Handle error
			return nil, err
		}
	}

	go func() {
		itemCount := 0
		for {
			item := <- out
			id, err := save(client,index,item)
			if err != nil {
				log.Printf("item save :error saving item %v : %v\n", item, err )
			}
			fmt.Printf("es save client count:%d id %v ,item : %v\n", itemCount,id, item)
			itemCount ++
		}
	}()
	return out,nil
}

func save(client *elastic.Client ,index string, item interface{} ) (id string,err error){
	//data := model.Member(item)
	data := item.(model.Member)
	IdString  := strconv.Itoa(data.MemberID)
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	exist, err := client.Exists().
		Index(index).
		Type(Type).
		Id(IdString).
		Do(ctx)
	if err != nil  {
		return  "", err
	}
	if exist {
		resp, err := client.Update().
			Index(index).
			Type(Type).
			Id(IdString).
			Doc(item).
			Do(ctx)
		if err != nil {
			return "", err
		}
		return resp.Id, nil
	} else {
		resp, err := client.Index().
			Index(index).
			Type(Type).
			Id(IdString).
			BodyJson(item).
			Do(ctx)
		if err != nil {
			return "", err
		}
		return resp.Id, nil
	}

}