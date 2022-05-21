package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"regexp"
	"strconv"
)

var CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contens []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contens,-1)
	result := engine.ParseResult{}
	limit := 1
	PageLimit := 6
	for _, m := range matches {
		if limit  == 0 {
			break
		}
		result.Items = append(result.Items,"City " +string(m[2]))
		for i :=1; i <= PageLimit; i ++ {
			url := string(m[1]) + "/" + strconv.Itoa(i)
			result.Requests = append(result.Requests,engine.Request{
				Type: "url",
				Url: url,
				ParserFuc: func(bytes []byte) engine.ParseResult {
					return ParseCityUserList(bytes, url)
				},
			})
		}
		limit --
	}

	return result
}
