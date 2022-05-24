package engine


type ParseFunc func(contents []byte, url string) ParseResult


type Request struct {
	Type string  // url, html
	Url  string
	Text string
	//ParserFuc ParseFunc
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items  []Item
}

type Item struct {
	Id int
	Url string
	Type string
	Payload interface{}
}

type FuncParser struct {
	 parser ParseFunc
	 funcName string
}

type Parser interface {
	Parser(contents []byte, url string) ParseResult
	Serialize() (funcName string, args interface{})
}

func (f *FuncParser) Parser(contents []byte, url string) ParseResult {
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (funcName string, args interface{}) {
	return f.funcName, nil
}

func NewFuncParser(p ParseFunc, funcName string) *FuncParser {
	return &FuncParser{
		parser: p,
		funcName: funcName,
	}
}


//func NilParser([]byte) ParseResult {
//	return ParseResult{}
//}

type NilParser struct {

}

func (f NilParser) Parser(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (f NilParser) Serialize() (_ string, _ interface{}) {
	return "NilParser", nil
}


