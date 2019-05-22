package engine

import (
	"IMOOC/crawler_distributed/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching url %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}

