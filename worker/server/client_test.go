package main

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/rpcsupport"
	"IMOOC/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := engine.SerializedRequest{
		"http://album.zhenai.com/u/1172388090",
		engine.SerializedParser{
			config.ParseProfile,
			"思忆",
		},
	}
	var result engine.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	}else{
		fmt.Println(result)
	}

}
