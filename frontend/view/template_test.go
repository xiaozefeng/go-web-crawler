package view

import (
	"bufio"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/frontend/model"
	"github.com/xiaozefeng/go-web-crawler/model/zhenai"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	var item = engine.Item{
		Id:"123",
		Url:"https://album.zhenai.com/u/1755238721",
		Type:"zhenai",
		Payload: zhenai.Profile{
			Name:          "非诚勿扰",
			Gender:        "男",
			Age:           "12",
			Height:        "170cm",
			Weight:        "150kg",
			Income:        "20000-30000",
			Marriage:      "未婚",
			Education:     "大专",
			Occupation:    "工程师",
			Hukou:         "上海",
			Constellation: "双鱼",
			House:         "已购房",
			Car:           "已购车",
			Avatar:        "https://photo.zastatic.com/images/photo/438810/1755238721/1729715236718796.jpg?scrop=1&crop=1&cpos=north&w=200&h=200",
		},
	}
	temp := template.Must(template.ParseFiles("./index.tmpl"))
	file, err := os.Create("./index.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	page := model.SearchResult{
		Hits:  10,
		Start: 0,
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = temp.Execute(w, page)
	defer w.Flush()
	if err != nil {
		panic(err)
	}

}
