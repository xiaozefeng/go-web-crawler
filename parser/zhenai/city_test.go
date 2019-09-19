package zhenai

import (
	"bufio"
	"fmt"
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveCityData(t *testing.T) {
	content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("./city_data.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, _ = writer.Write(content)
	defer writer.Flush()
}

func TestParseCityWithGoQuery(t *testing.T) {
	bytes, err := ioutil.ReadFile("./city_data.html")
	if err != nil {
		panic(err)
	}
	parseResult := ParseCityWithGoQuery(bytes)
	for _, v := range parseResult.Requests {
		fmt.Printf("%s\n", v.Url)
	}

	fmt.Println(parseResult.Items)
}
