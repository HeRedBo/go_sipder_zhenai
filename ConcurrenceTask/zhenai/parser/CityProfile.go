package parser

import (
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/zhenai/model"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func ParseCityProfile (contens []byte) engine.ParseResult {
	ioReader := bytes.NewReader(contens)
	doc, err := goquery.NewDocumentFromReader(ioReader)
	if err != nil {
		log.Fatal(err)
	}

	photo,_  := doc.Find(".photo >a>img").Attr("src")
	nickName := doc.Find(".content a").Eq(0).Text()
	gender 	 := doc.Find(".content tr").Eq(1).Find("td").Eq(0).Text()
	position := doc.Find(".content tr").Eq(1).Find("td").Eq(1).Text()
	age 	 := doc.Find(".content tr").Eq(2).Find("td").Eq(0).Text()
	salary 	 := doc.Find(".content tr").Eq(2).Find("td").Eq(1).Text()
	marriage := doc.Find(".content tr").Eq(3).Find("td").Eq(0).Text()
	height := doc.Find(".content tr").Eq(3).Find("td").Eq(1).Text()
	introduce := doc.Find(".introduce").Text()

	gender2  := strings.Replace(gender,"性别：","",1)
	position2  := strings.Replace(position,"居住地：","",1)
	age2  := strings.Replace(age,"年龄：","",1)
	salary2  := strings.Replace(salary,"月   薪：","",1)
	marriage2  := strings.Replace(marriage,"婚况：","",1)
	height2  := strings.Replace(height,"身   高：","",1)

	CityProfile  := model.CityProfile{}
	CityProfile.Name = nickName
	CityProfile.Photo = photo
	CityProfile.Gender = gender2
	CityProfile.Place = position2
	CityProfile.Age = age2
	CityProfile.Income = salary2
	CityProfile.Marriage = marriage2
	CityProfile.Height = height2
	CityProfile.Introduce = introduce

	result := engine.ParseResult{
		Items: []interface{}{CityProfile},
	}
	return result










}
