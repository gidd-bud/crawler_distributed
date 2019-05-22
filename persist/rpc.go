package persist

import (
	"IMOOC/crawler_distributed/engine"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

type ItemSaverService struct {
	Client *elasticsearch.Client
	Index string

}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
		log.Printf("Item %v saved.\n", item)
	}else{
		log.Printf("Error saving item %v: %v.\n", item, err)
	}
	return err
}
