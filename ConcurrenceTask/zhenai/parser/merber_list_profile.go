package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/zhenai/model"
	"encoding/json"
	"fmt"
	"strconv"
)

func ParseMemberListProfile (contens []byte, url string) engine.ParseResult {
	data := model.CityDetailList{}
	result := engine.ParseResult{}
	if err := json.Unmarshal(contens,&data); err == nil {
		member_list := data.MemberListData.MemberList
		
		for _,item := range member_list {
			member := model.Member(item)
			item := engine.Item{
				Id:      member.MemberID,
				Url:     "http://album.zhenai.com/u/" +   strconv.Itoa(member.MemberID),
				Type:    "zhengai",
				Payload: item,
			}
			result.Items = append(result.Items,item)
		}
	} else {
		fmt.Println("error:" + url)
		panic(err)
	}
	return result
}
