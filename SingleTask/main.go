package main

import (
	"GoSpider/SingleTask/engine"
	"GoSpider/SingleTask/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Type : "url",
		Url: "http://www.zhenai.com/zhenghun",
		ParserFuc: parser.ParseCityList,
	})
}




