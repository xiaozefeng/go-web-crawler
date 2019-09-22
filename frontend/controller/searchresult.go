package controller

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/xiaozefeng/go-web-crawler/distributed/config"
	"github.com/xiaozefeng/go-web-crawler/engine"
	"github.com/xiaozefeng/go-web-crawler/frontend/model"
	"github.com/xiaozefeng/go-web-crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(t string) (SearchResultHandler, error) {
	var r SearchResultHandler
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return r, err
	}
	r = SearchResultHandler{
		view:   view.CreateSearchResultView(t),
		client: client,
	}
	return r, nil
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	page, err := s.getPageData(rewriteQueryString(q), from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	page.Query = q
	err = s.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s SearchResultHandler) getPageData(q string, from int) (model.SearchResult, error) {
	var r = model.SearchResult{}
	resp, err := s.client.
		Search(config.ElasticSearchIndex).
		//Type("zhenai").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		return r, err
	}
	r.Hits = resp.TotalHits()
	r.Start = from
	r.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	r.PrevFrom = r.Start - len(r.Items)
	r.NextFrom = r.Start + len(r.Items)
	return r, nil

}

func rewriteQueryString(s string) string {
	var re = regexp.MustCompile(`([A-Z][a-z]*)`)
	return re.ReplaceAllString(s, "Payload.$1:")
}
