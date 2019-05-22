package persist

import (
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/model"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"testing"
)

func TestSave(t *testing.T) {

	expectedItem := engine.Item{
		Url:  "http://album.zhenai.com/u/1172388090",
		Type: "zhenai",
		Id:   "1172388090",
		Payload: model.Profile{
			"1172388090",
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
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(esClient, index, expectedItem)
	if err != nil {
		panic(err)
	}

	// TODO: try to start up elastic search
	// here using docker go client.
	getRequest := esapi.GetRequest{
		Index:        index,
		DocumentType: expectedItem.Type,
		DocumentID:   expectedItem.Id,
	}
	response, err := getRequest.Do(context.Background(), esClient)

	if err != nil {
		panic(err)
	}
	t.Logf("%+v", response)

	// refer to the github
	var r map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		panic(err)
	}

	var actualProfile model.Profile
	err = actualProfile.FromJsonObj(r["_source"].(map[string]interface{})["Payload"])

	if actualProfile != expectedItem.Payload {
		t.Errorf("Got %v; Expected %v.\n", actualProfile, expectedItem.Payload)
	}

}
