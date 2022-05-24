package worker

import (
	"GoSpider/ConcurrenceTask/crawler_distributed/config"
	"GoSpider/ConcurrenceTask/engine"
	"GoSpider/ConcurrenceTask/zhenai/parser"
	"fmt"
	"log"
)

type SerializedParser struct {
	FuncName string
	Args interface{}
}

type Request struct {
	Url string
	Type  string
	Text string
	Parser SerializedParser
}

type ParserResult struct {
	 Items  []engine.Item
	Requests []Request
}

func SerializeRuquest(r engine.Request) Request {
	funcName, args := r.Parser.Serialize()
	req := Request{
		r.Url,
		r.Type,
		r.Text,
		SerializedParser{
			funcName,
			args,
		},
	}
	return req
}

func SerializeParserResult(result engine.ParseResult) ParserResult {
	parseResult := ParserResult{
		Items: result.Items,
	}
	for _, req := range result.Requests {
		serialReq := SerializeRuquest(req)
		parseResult.Requests =  append(parseResult.Requests, serialReq)
	}

	return parseResult
}


func DeserializeRequest(req Request) (engine.Request,error) {
	deserializeParser , err := DeserializeParser(req.Parser)
	if err !=nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Type: req.Type,
		Url: req.Url,
		Text: req.Text,
		Parser: deserializeParser,
	},nil

}

func DeserializeParser(p SerializedParser) (engine.Parser,error) {
	switch p.FuncName {
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),nil
	case config.ParseCityUserList:
		return engine.NewFuncParser(parser.ParseCityUserList,config.ParseCityUserList),nil
	case config.ParseMemberListProfile:
		return engine.NewFuncParser(parser.ParseMemberListProfile,config.ParseMemberListProfile),nil
	default:
		return nil, fmt.Errorf("unknown parser func name %s", p.FuncName)
	}
}


func DeserializeResult(r ParserResult) (engine.ParseResult,error) {

	desParserResult := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request,err  := DeserializeRequest(req)
		if err != nil {
			log.Printf("DeserializeRequest %+v Error:%+v\n", request, err)
			continue
		}
		desParserResult.Requests = append(desParserResult.Requests,request)
	}

	return desParserResult,nil
}