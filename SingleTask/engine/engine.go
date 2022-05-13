package engine

import (
	"GoSpider/SingleTask/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for  _, r := range seeds {
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		var body []byte
		log.Printf("Fetchin g type %s: Url: %s",  r.Type, r.Url)
		if r.Type == "url" {
			var err error
			body ,err  = fetcher.Fetch(r.Url)
			if err != nil {
				log.Printf("Fetcher: error " + " fetching url %s : %s" , r.Url, err)
				continue
			}
		} else if r.Type == "html"  {
			var data []byte = []byte(r.Text)
			body  = data
		}
		parseResult := r.ParserFuc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}

}
