package parser

import (
	"GoSpider/SingleTask/engine"
	"GoSpider/SingleTask/zhenai/model"
	"encoding/json"
	"fmt"
)

func ParseMemberListProfile (contens []byte, url string) engine.ParseResult {
	data := model.CityDetailList{}
	result := engine.ParseResult{}
	if err := json.Unmarshal(contens,&data); err == nil {
		member_list := data.MemberListData.MemberList
		for _,item := range member_list {
			member := model.Member(item)
			result.Items = append(result.Items, member)
		}
	} else {
		fmt.Println("error:" + url)
		panic(err)
	}
	return result
}
