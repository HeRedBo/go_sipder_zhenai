package controller

import (
	"GoSpider/ConcurrenceTask/frontend/model"
	"GoSpider/ConcurrenceTask/frontend/view"
	model2 "GoSpider/ConcurrenceTask/zhenai/model"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	cient *elastic.Client
}


func CreateSearchResultHandler(template string ) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view : view.CreateSearchResultView(template),
		cient: client,
	}
}

// localhost:8888/search?q=男 以购房&from=20
func (s SearchResultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	q := strings.TrimSpace(request.FormValue("q"))
	from , err := strconv.Atoi(request.FormValue("from"))
	if err !=nil {
		from = 0
	}
	size , err := strconv.Atoi(request.FormValue("size"))
	if err !=nil {
		size = 10
	}
	p ,err := strconv.Atoi(request.FormValue("p"))
	if err !=nil {
		p = 1
	}




	query := rewriteQueryString(q)
	var page model.SearchResult
	page, err = s.GetSearchResult(query, p,size)
	if err != nil {
		http.Error(writer,err.Error(), http.StatusBadRequest)
	}
	page.Query = q
	err = s.view.Render(writer, page)
	if err != nil {
		http.Error(writer,err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(writer,"q=%s, from=%d",q,from)
}

func (s SearchResultHandler) GetSearchResult(q string, p int,size int) (model.SearchResult,error) {
	var result model.SearchResult

	from := ( p -1) * size

	query_service := s.cient.Search("dating_prifile").
		From(from).
		Size(size)

	if q != "" {
		query_service.Query(elastic.NewQueryStringQuery(q))
	}
	resp, err := query_service.Do(context.Background())
	if err != nil {
		return result, err
	}
	fmt.Println(q)
	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(model2.Member{}))
	result.CurrentPage = p
	result.TotalPage = int64(math.Ceil(float64(result.Hits /int64(size))))

	//result.TotalPage = result /
	return result, nil
}

func rewriteQueryString(q string) string {

	q = strings.Replace(q,"男","Sex:0",-1)
	q = strings.Replace(q,"男士","Sex:0",-1)
	q = strings.Replace(q,"女士","Sex:1",-1)
	q = strings.Replace(q,"女士","Sex:1",-1)
	return q
}




