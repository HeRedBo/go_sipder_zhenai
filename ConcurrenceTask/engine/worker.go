package engine

import (
	"GoSpider/ConcurrenceTask/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	var body []byte
	log.Printf("Fetching  type %s: Url: %s",  r.Type, r.Url)
	if r.Type == "url" {
		var err error
		body ,err  = fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error " + " fetching url %s : %s" , r.Url, err)
			return ParseResult{}, err
		}
	} else if r.Type == "json" || r.Type == "html" {
		var data = []byte(r.Text)
		body  = data
	}
	//parseResult := r.ParserFuc(body,r.Url)
	parseResult := r.Parser.Parser(body,r.Url)
	return parseResult, nil
}