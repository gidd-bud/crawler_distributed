package main

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/persist"
	"IMOOC/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var port = flag.Int("port",0,"The port for itemsaver server to listen on.")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Port must be specified.")
		return
	}
	log.Printf("Itemsaver server listening on port %d.", *port)

	log.Fatal(serverRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serverRpc(host, index string) error {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		esClient,
		index,
	})
}