package zhenai

import (
	"bufio"
	"fmt"
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"io/ioutil"
	"os"
	"strings"
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

func TestParseCityWithGoQuery2(t *testing.T) {
	bytes, err := ioutil.ReadFile("./city_data.html")
	if err != nil {
		panic(err)
	}
	matches := cityURLRe.FindAllStringSubmatch(string(bytes), -1)
	for _, m := range matches {
		fmt.Println(m[1])
	}
}

func TestGetIdFromURL(t *testing.T) {
	var s = "https://album.zhenai.com/u/1770372887"
	if index := strings.LastIndex(s, `/`); index!=-1{
		s = s[index+1:]
	}
	fmt.Println("s",s)

}


