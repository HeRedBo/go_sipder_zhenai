package main

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/crawler_distributed/rpcsupport"
	model2 "GoSpider/ConcurrenceTask/zhenai/model"
	"fmt"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	host := fmt.Sprintf(":%d", config.ItemSaverServicePort)
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client , err := rpcsupport.NewCient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := model2.Member{
		Age : 49,
		AvatarURL : "https://photo.zastatic.com/images/photo/482882/1931527315/68699857905725727.jpg",
		Car : "未填写",
		Children : "有孩子但不在身边",
		Constellation : "魔羯座(12.22-01.19)",
		Education : "高中及以下",
		Height : 156,
		House : "住在单位宿舍",
		IntroduceContent : "孤独的我，只想找个温柔，体贴，有责任心的相互理解，尊重，包容。",
		IsRecommend : 0,
		LastModTime : "2022-01-15 16:02:36",
		Marriage : "离异",
		MemberID : 1931527315,
		NickName : "岁月静好",
		ObjectAge : "48-52岁",
		ObjectHight : "174-174cm",
		ObjectMarriage : "丧偶",
		ObjectSalary : "5000元以上",
		Occupation : "保险",
		Salary : "3001-5000元",
		Sex : 1,
		WorkCity : "四川阿坝",
	}

	result := ""
	err = client.Call(config.CrawlerServiceRpc,item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err %s", result, err)
	}


}
