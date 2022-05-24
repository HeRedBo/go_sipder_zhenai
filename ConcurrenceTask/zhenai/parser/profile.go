package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/zhenai/model"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
)

const ProfileRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const ProfileRe2 = `<div data-v-8b1eac0c="" class="m-btn purple">未婚</div>`

func ParseProfile (contens []byte, _ string ) engine.ParseResult {
	ioReader := bytes.NewReader(contens)
	doc, err := goquery.NewDocumentFromReader(ioReader)
	if err != nil {
		log.Fatal(err)
	}
	userInfoMap := make(map[int]string)
	doc.Find(".m-content-box >.purple-btns .purple").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		userInfoMap[i] = content
	})
	nickName := doc.Find(".info h1.nickName").Text()
	Id := doc.Find(".info .id").Text()
	id ,err  := strconv.Atoi(strings.Replace(Id,"ID：","",1))
	if err != nil {
		panic(err)
	}
	position := doc.Find(".seoFooter >p>span>a").Eq(2).Text()
	var gender string
	if strings.Contains(position, "女士") {
		gender = `女士`
	} else {
		gender = `男士`
	}
	profile := model.Profile{}
	profile.ID = id
	profile.Name = nickName
	//profile.Name = name
	profile.Gender = gender
	profile.Marriage = userInfoMap[0]
	profile.Age = userInfoMap[1]
	profile.Xinzuo = userInfoMap[2]
	profile.Height = userInfoMap[3]
	profile.Weight = userInfoMap[4]
	profile.Income = userInfoMap[6]
	profile.Education = userInfoMap[8]
	//result := engine.ParseResult{}
	//result.Items =
	//result := engine.ParseResult{
	//	Items: []interface{}{profile},
	//}
	result := engine.ParseResult{}
	result.Items = append(result.Items,engine.Item{Payload: profile})

	//result := engine.ParseResult{
	//	Items: []engine.Item{}{ Payload : profile},
	//}
	return result
}