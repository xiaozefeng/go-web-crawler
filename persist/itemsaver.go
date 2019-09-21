package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	r := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-r
			itemCount++
			log.Printf("Got item:%v, itemCount:%d\n", item, itemCount)
			err := Save(client, index, item)
			if err != nil {
				log.Println("item saver, Save item err:", err)
				continue
			}
		}
	}()
	return r, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())
	return err
}
