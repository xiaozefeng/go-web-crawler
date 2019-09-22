package client

import (
	"github.com/xiaozefeng/go-web-crawler/distributed/config"
	"github.com/xiaozefeng/go-web-crawler/distributed/rpcsupport"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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
			// call rpc
			var result string
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Println("item saver, Save item err:", err)
				continue
			}
		}
	}()
	return r, nil
}
