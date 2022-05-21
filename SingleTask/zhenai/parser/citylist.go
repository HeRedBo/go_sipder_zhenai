package parser

import (
	"GoSpider/SingleTask/engine"
	"regexp"
	"strconv"
)

var CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contens []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contens,-1)
	result := engine.ParseResult{}
	//limit := 1
	page_limit := 6 // 页面限制数据为 6 暂时取 6页 数据
	for _, m := range matches {
		result.Items = append(result.Items,"City " +string(m[2]))
		// 追加页面翻页数据请求
		for i :=1; i <= page_limit; i ++ {
			url := string(m[1]) + "/" + strconv.Itoa(i)
			result.Requests = append(result.Requests,engine.Request{
				Type: "url",
				Url: url,
				ParserFuc: func(bytes []byte) engine.ParseResult {
					return ParseCityUserList(bytes, url)
				},
			})
		}

		//limit --
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
