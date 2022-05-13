package engine

import (
	"GoSpider/ConcurrenceTask/fetcher"
	"log"
)

type SimpleEngine struct {}


func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for  _, r := range seeds {
		requests = append(requests,r)
	}
	for len(requests) > 0 {

		r := requests[0]
		requests = requests[1:]

		parseResult ,err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	var body []byte
	log.Printf("Fetching  type %s: Url: %s",  r.Type, r.Url)
	if r.Type == "url" {
		var err error
		body ,err  = fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error " + " fetching url %s : %s" , r.Url, err)
			return ParseResult{}, err
		}
	} else if r.Type == "html"  {
		var data []byte = []byte(r.Text)
		body  = data
	}
	parseResult := r.ParserFuc(body)
	return parseResult, nil
}