package client

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/zhenai.com/parser"
	"net/rpc"
)

func CreateProcessor(clientchan chan *rpc.Client) engine.Processor {
	return func(request engine.Request) (result engine.ParseResult, e error) {
		sReq := parser.SerializeRequest(request)
		var sResult engine.SerializedParseResult

		c := <-clientchan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}else{
			return parser.DeserializeParserResult(sResult), nil
		}
	}
}
