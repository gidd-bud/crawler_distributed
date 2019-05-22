package engine

import "encoding/json"

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() SerializedParser
}

type Request struct {
	Url 	string
	Parser 	Parser
}

type ParseResult struct {
	Requests []Request
	Items	 []Item
}

type Item struct {
	Url	 		string
	Type 		string
	Id 			string
	Payload 	interface{}
}

func (i *Item)FromJsonObj(o interface{}) error {
	bytes, err := json.Marshal(o)
	if err == nil {
		err = json.Unmarshal(bytes, i)
	}
	return err
}

type NilParser struct {}

func (NilParser) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() SerializedParser {
	return SerializedParser{
		"NilParser",
		nil,
	}
}

/// for serialized
type SerializedParser struct {
	Name string
	Args interface{}
}

type SerializedRequest struct {
	Url string
	Parser SerializedParser
}

type SerializedParseResult struct {
	Requests 	[]SerializedRequest
	Items 		[]Item
}


