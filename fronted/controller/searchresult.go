package controller

import (
	"IMOOC/crawler_distributed/config"
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/fronted/model"
	"IMOOC/crawler_distributed/fronted/view"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elasticsearch.Client
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page model.SearchResult
	page, err = h.getSearchResult(q, &from)
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from *int) (model.SearchResult, error) {
	var result model.SearchResult
	queryRequest := esapi.SearchRequest{
		Index:        []string{config.ElasticIndex},
		DocumentType: []string{"zhenai"},
		Query:        RewriteQueryString(q),
		From:         from,
	}
	response, err := queryRequest.Do(context.Background(), h.client)
	if err != nil {
		panic(err)
	}

	// refer to the github
	var r map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		panic(err)
	}

	tmpHits := r["hits"].(map[string]interface{})
	hits := tmpHits["hits"].([]interface{})
	result.Hits = len(hits)
	result.Start = *from
	result.Query = q

	var tmpItem engine.Item
	for _, hit := range hits {
		err = tmpItem.FromJsonObj(hit.(map[string]interface{})["_source"])

		result.Items = append(result.Items, tmpItem)
	}

	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view.CreateSearchResultView(template),
		esClient,
	}
}

func RewriteQueryString(q string) string {
	compile := regexp.MustCompile(`([A-Za-z]*):`)
	return compile.ReplaceAllString(q, "Payload.$1:")
}