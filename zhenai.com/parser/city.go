package parser

import (
	"IMOOC/crawler_distributed/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+")`)
)
func ParseCity(contents []byte, _ string) engine.ParseResult {
	allSubmatch := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, sub := range allSubmatch{
		result.Requests = append(result.Requests, engine.Request{string(sub[1]), NewProfileParser(string(sub[2]))})
	}

	allSubmatch = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, sub := range allSubmatch{
		result.Requests = append(result.Requests, engine.Request{
			string(sub[1]),
			NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}

