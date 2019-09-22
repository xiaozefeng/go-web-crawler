package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"github.com/xiaozefeng/go-web-crawler/distributed/config"
	"github.com/xiaozefeng/go-web-crawler/distributed/persist"
	"github.com/xiaozefeng/go-web-crawler/distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort),
		config.ElasticSearchIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	err = rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	if err != nil {
		return err
	}
	return nil
}
