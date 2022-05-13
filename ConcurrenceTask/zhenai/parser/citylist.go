package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"regexp"
)

var CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contens []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contens,-1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items,"City " +string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Type: "url",
			Url:string(m[1]),
			//ParserFuc: engine.NilParser,
			ParserFuc: ParseCity,
		})
	}
	return result
}