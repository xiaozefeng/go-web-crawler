package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id string
	Url string
	Type string
	Payload interface{}
}

func NiParser([]byte) ParseResult{
	return ParseResult{}
}
