package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/xiaozefeng/go-web-crawler/distributed/config"
	"testing"
)

func TestElastic(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	_, err = client.Search(config.ElasticSearchIndex).Query(elastic.NewQueryStringQuery("123asfafafasfa")).
		From(10).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
}

