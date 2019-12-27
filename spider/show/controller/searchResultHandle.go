package controller

import (
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"regexp"
	"spider/engine"
	"spider/show/model"
	"spider/show/view"
	"strconv"
	"strings"
)

func CreateSearchResultHandle(template string) (SearchResultHandle){
	client, err := elastic.NewClient(
		elastic.SetURL("http://www.vowcloud.cn:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	return SearchResultHandle{
		view: view.CreateSearchResultView(template),
		client: client,
	}

}

type SearchResultHandle struct {
	view view.SearchResultView
	client *elastic.Client
}

func (s SearchResultHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := s.getSearchResult(q, from)
	if err != nil {
		panic(err)
	}

	err = s.view.Render(w, page)
	if err != nil {
		panic(err)
	}
}

func (s SearchResultHandle) getSearchResult (q string,
	from int) (model.SearchResultData, error){
	var result model.SearchResultData
	resp, err := s.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQuery(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result,err
	}

	result.Hits = int(resp.TotalHits())
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.Count = len(result.Items)
	result.PrevPage = (from - len(result.Items))
	result.NextPage = (from + len(result.Items))
	result.Query = q

	return result, nil
}

func rewriteQuery(q string) string  {
	regx := regexp.MustCompile(`([A-Z][a-z]*):`)
	return regx.ReplaceAllString(q, "Payload.$1:")
}
