package zhenai

import (
	"bufio"
	"fmt"
	"github.com/xiaozefeng/go-web-crawler/fetcher"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveProfile(t *testing.T) {
	content, err := fetcher.Fetch("https://album.zhenai.com/u/108208979")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("./profile_data.html")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	_, _ = writer.Write(content)
	defer writer.Flush()
}

func TestParseProfile(t *testing.T) {
	content, err := ioutil.ReadFile("./profile_data.html")
	if err != nil {
		panic(err)
	}
	parseResult:= ParseProfile(content,"","", "")
	fmt.Printf("%#v",parseResult.Items[0])
}



