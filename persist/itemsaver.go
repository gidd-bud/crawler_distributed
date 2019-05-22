package persist

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/rpcsupport"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <- out
			log.Printf("Item Saver: got item #%d: %v\n", itemCount, item)
			itemCount++

			// call RPC
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v.\n", item, err)
			}
		}
	}()
	return out, nil
}

func save(esClient *elasticsearch.Client, index string, item engine.Item) error {
	// Must turn off Sniff in docker. but i dont know how to set in my project.

	if item.Type == "" {
		return errors.New("must supply type!")
	}
	itemStr, _ := json.Marshal(item)
	request := esapi.IndexRequest{
		Index:        	index,
		DocumentType: 	item.Type,
		//DocumentID:	 	item.Id,
		Body:         	strings.NewReader(string(itemStr)),
	}
	if item.Id != "" {
		request.DocumentID = item.Id
	}
	response, err := request.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// refer to the github
	var r map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		return err
	}
	return nil
}
