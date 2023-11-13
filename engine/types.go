package engine

type Request struct {
	Url        string
	ParserFunc func(string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
