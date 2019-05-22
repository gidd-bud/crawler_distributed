package model

import (
	"IMOOC/crawler_distributed/engine"
)

type SearchResult struct {
	Hits     int
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []engine.Item
}
