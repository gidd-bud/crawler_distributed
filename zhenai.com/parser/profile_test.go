package parser

import (
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile("http://album.zhenai.com/u/1172388090", contents, "思忆")

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

	if len(result.Items) != 1 {
		t.Errorf("Result should have %d items, but had %d.", 1, len(result.Items))
	}

	actualItem := result.Items[0]
	if actualItem != expectedItem {
		t.Errorf("Expected profile: %v, but was %v.", expectedItem.Payload, actualItem.Payload)
	}

	

}