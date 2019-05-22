package main

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/model"
	"IMOOC/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// Start ItemSaverServer
	go serverRpc(host, "test1")
	time.Sleep(time.Second)
	// Start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	// Call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1172388090",
		Type: "zhenai",
		Id:   "1172388090",
		Payload: model.Profile{
			1172388090,
			"思忆",
			24,
			"女士",
			165,
			"大学本科",
			"上海",
			"未婚",
			"20001-50000元",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("Result: %v; Error: %v.\n",result, err)
	}
}
