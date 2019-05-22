package parser

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/engine"
	"errors"
	"fmt"
	"log"
)

type ParserFunc func(contents []byte, url string) engine.ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) engine.ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() engine.SerializedParser {
	return engine.SerializedParser{
		f.name, nil,
	}
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: parser,
		name:   name,
	}
}

type ProfileParser struct {
	username string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(url, contents, p.username)
}

func (p *ProfileParser) Serialize() engine.SerializedParser {
	return engine.SerializedParser{
		"ParseProfile", p.username,
	}
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		name,
	}
}


// for types conversion
func SerializeRequest(r engine.Request) engine.SerializedRequest {
	return engine.SerializedRequest{
		r.Url,
		r.Parser.Serialize(),
	}
}

func SerializeParseRequest(r engine.ParseResult) engine.SerializedParseResult {
	result := engine.SerializedParseResult{
		Items: r.Items,
	}
	for _, request := range r.Requests{
		result.Requests = append(result.Requests, SerializeRequest(request))
	}
	return result
}

func deserializeParser(p engine.SerializedParser) (engine.Parser,error) {
	switch p.Name {
	case config.ParseCityList:
		return NewFuncParser(ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return NewFuncParser(ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if username, ok := p.Args.(string); ok{
			return NewProfileParser(username), nil
		}else{
			return nil, fmt.Errorf("Invalid arg: %v.\n", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("Unknow parser name.")
	}
}

func DeserializeRequest(r engine.SerializedRequest) (engine.Request, error) {
	if p, err := deserializeParser(r.Parser); err!=nil{
		return engine.Request{}, err
	}else{
		return engine.Request{
			r.Url,
			p,
		}, nil
	}
}

func DeserializeParserResult(r engine.SerializedParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items:r.Items,
	}
	for _, request := range r.Requests{
		if req, err := DeserializeRequest(request); err!=nil{
			log.Printf("Error deseralizeing request %v.\n", request)
		}else{
			result.Requests = append(result.Requests, req)
		}
	}
	return result
}
