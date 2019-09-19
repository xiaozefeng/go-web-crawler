package parser

import (
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("./citylist_data.html")
	if err != nil {
		panic(err)
	}
	parseResult := ParseCityList(content)
	var resultSize = 470

	if len(parseResult.Requests) != resultSize {
		t.Errorf("Expect url: %d ,but got url: %d", resultSize, len(parseResult.Requests))
	}

	if len(parseResult.Items) != resultSize {
		t.Errorf("Expect item %d, buit got %d", resultSize, len(parseResult.Items))
	}
}

func TestSaveCityList(t *testing.T) {
	content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("./citylist_data.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, _ = file.Write(content)
}
