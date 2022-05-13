package engine


type Request struct {
	Type string  // url, html
	Url  string
	Text string
	ParserFuc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items  []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}