package worker

import (
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/zhenai.com/parser"
)

type CrawlService struct {}

func (CrawlService) Process(req engine.SerializedRequest, result *engine.SerializedParseResult) error {
	request, err := parser.DeserializeRequest(req)
	if err != nil {
		return err
	}
	parseResult, err := engine.Worker(request)
	if err != nil {
		return err
	}
	*result = parser.SerializeParseRequest(parseResult)
	return nil
}

