package parser

import (
	"IMOOC/crawler_distributed/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	all := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, sub := range all{
		result.Requests = append(result.Requests, engine.Request{string(sub[1]), NewFuncParser(ParseCity,"ParseCity")})
	}
	return result
}
