package main

import (
	"IMOOC/crawler_distributed/rpcsupport"
	"IMOOC/crawler_distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port",0,"The port for worker server to listen on.")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Port must be specified.")
		return
	}
	log.Printf("Worker server listening on port %d.", *port)

	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{},
		))
}
