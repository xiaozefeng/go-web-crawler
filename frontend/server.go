package main

import (
	"github.com/xiaozefeng/go-web-crawler/frontend/controller"
	"log"
	"net/http"
)

func main() {
	handler, err := controller.CreateSearchResultHandler("./view/index.tmpl")
	if err != nil {
		panic(err)
	}
	http.Handle("/search", handler)
	log.Println("http server incoming at address 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
