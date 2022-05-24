package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"regexp"
	"strconv"
)

var CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contens []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contens,-1)
	result := engine.ParseResult{}
	limit := 100
	PageLimit := 6

	for _, m := range matches {
		if limit  == 0 {
			break
		}
		name := "City " + string(m[2])
		item := engine.Item{
			Payload: name,
		}
		result.Items = append(result.Items,item)
		for i :=1; i <= PageLimit; i ++ {
			url := string(m[1]) + "/" + strconv.Itoa(i)
			result.Requests = append(result.Requests,engine.Request{
				Type: "url",
				Url: url,
				//ParserFuc: func(bytes []byte, url string ) engine.ParseResult {
				//	return ParseCityUserList(bytes, url)
				//},
				//ParserFuc : ParseCityUserList,
				Parser: engine.NewFuncParser(ParseCityUserList,"ParseCityUserList"),
			})
		}
		limit --
	}
	return result
}

