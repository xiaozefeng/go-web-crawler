package main

import (
	"github.com/olivere/elastic"
	"github.com/xiaozefeng/go-web-crawler/distributed/persist"
	"github.com/xiaozefeng/go-web-crawler/distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(serveRpc())
}

func serveRpc() error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	err = rpcsupport.ServeRpc(":1234", &persist.ItemSaverService{
		Client: client,
		Index:  "dating_profile",
	})
	if err != nil {
		return err
	}
}
