package main

import (
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/persist"
	"IMOOC/crawler_distributed/rpcsupport"
	"IMOOC/crawler_distributed/scheduler"
	"IMOOC/crawler_distributed/worker/client"
	"IMOOC/crawler_distributed/zhenai.com/parser"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemSaverHost", "", "itemsaver host")
	workerHosts   = flag.String("workerHosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	if *itemSaverHost == "" || *workerHosts == "" {
		fmt.Println("Port must be specified.")
		return
	}

	itemSaver, err := persist.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := creatClientPool(strings.Split(*workerHosts, ","))
	processor := client.CreateProcessor(pool)
	e := engine.ConcurrentEngine{&scheduler.QueuedScheduler{}, 100, itemSaver, processor}

	e.Run(engine.Request{"http://www.zhenai.com/zhenghun/", parser.NewFuncParser(parser.ParseCityList, "ParseCityList")})
}

func creatClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		c, err := rpcsupport.NewClient(host)
		if err != nil {
			log.Printf("Error connecting to %s: %v.\n", host, err)
		} else {
			clients = append(clients, c)
			log.Printf("Connected to %s.\n", host)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}

// TODO: 该分布式版本保存的数据格式错位，原crawler版本无此问题，需排查。
