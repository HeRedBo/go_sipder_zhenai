package parser

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/engine"
	"regexp"
	"strings"
)


var InstallRe = `window.__INITIAL_STATE__=(.*?)</script>`
func ParseCityUserList(contens []byte, url string) engine.ParseResult  {
	result := engine.ParseResult{}
	matches := regexp.MustCompile(InstallRe).FindAllSubmatch(contens,-1)
	for _,m := range matches {
		json_str := string(m[1])
		Text := strings.Replace(json_str,";(function(){var s;(s=document.currentScript||document.scripts[document.scripts.length-1]).parentNode.removeChild(s);}());","",1)
		result.Requests= append(result.Requests, engine.Request{
			Type: "json",
			Url: "detail:" + url,
			Text: Text,
			//ParserFuc: func(bytes []byte) engine.ParseResult {
			//	return MenberListProfileParser(bytes,url)
			//},
			//ParserFuc:ParseMemberListProfile,
			Parser: engine.NewFuncParser(ParseMemberListProfile,config.ParseMemberListProfile),
		})
	}
	return result
}

//func MenberListProfileParser(url string) engine.ParseFunc  {
//	return func(contens []byte) engine.ParseResult {
//		return ParseMemberListProfile(contens,url)
//	}
//}