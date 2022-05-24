package worker

import (
	"GoSpider/ConcurrenceTask/engine"
)

type CrawlerService struct {}

func (c *CrawlerService) Process(req Request,result *ParserResult) error {

	request, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	parseResult , err := engine.Worker(request)
	if err != nil {
		return err
	}
	*result = SerializeParserResult(parseResult)
	return nil
}