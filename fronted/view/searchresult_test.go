package view

import (
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/fronted/model"
	common "IMOOC/crawler_distributed/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, _ := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	testItem := engine.Item{
		Url:  "http://album.zhenai.com/u/1172388090",
		Type: "zhenai",
		Id:   "1172388090",
		Payload: common.Profile{
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

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, testItem)
	}
	err := view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
