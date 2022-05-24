package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func ParseCity (contens []byte, _ string) engine.ParseResult {
	ioReader := bytes.NewReader(contens)
	doc, err := goquery.NewDocumentFromReader(ioReader)
	if err != nil {
		log.Fatal(err)
	}
	result := engine.ParseResult{}
	doc.Find(".g-list .list-item").Each(func(i int, s *goquery.Selection) {
		nickName := s.Find(".content a").Eq(0).Text()
		href,_  := s.Find(".photo >a").Attr("href")
		content, _ := s.Html()
		name  := "User "  + nickName
		item := engine.Item{
			Payload: name,
		}
		result.Items = append(result.Items,item)
		result.Requests = append(result.Requests,engine.Request{
			Type: "html",
			Url : href,
			Text: content,
			//ParserFuc: ParseCityProfile,
			Parser: engine.NewFuncParser(ParseCityProfile,"ParseCityProfile"),
		})
	})
	return result
}