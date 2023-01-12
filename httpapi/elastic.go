package httpapi

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/oipwg/oip/datastore"
	"github.com/olivere/elastic/v7"
)

func GenerateNextAfter(hit *elastic.SearchHit) string {
	b, _ := json.Marshal(hit.Sort)
	return url.QueryEscape(string(b))
}

func ExtractSources(results *elastic.SearchResult) ([]json.RawMessage, string) {
	sources := make([]json.RawMessage, len(results.Hits.Hits))
	nextAfter := ""
	for k, v := range results.Hits.Hits {
		sources[k] = v.Source
		if k == len(results.Hits.Hits)-1 {
			nextAfter = GenerateNextAfter(v)
		}
	}
	return sources, nextAfter
}

func BuildCommonSearchService(ctx context.Context, indexNames []string, query elastic.Query, sorts []elastic.SortInfo, fsc *elastic.FetchSourceContext) *elastic.SearchService {
	var indices = make([]string, 0, len(indexNames))
	for _, index := range indexNames {
		indices = append(indices, datastore.Index(index))
	}

	searchService := datastore.Client().
		Search(indices...).
		TrackTotalHits(true).
		Query(query)

	size := GetSizeFromContext(ctx)
	searchService = searchService.Size(size)

	nSorts := GetSortInfoFromContext(ctx)
	nSorts = append(nSorts, sorts...)

	for _, v := range nSorts {
		searchService = searchService.SortWithInfo(v)
	}

	searchAfter := GetSearchAfterFromContext(ctx)
	if searchAfter != nil {
		searchService = searchService.SearchAfter(searchAfter...)
	}

	from := GetFromFromContext(ctx)
	if from != 0 && searchAfter == nil {
		searchService = searchService.From(from)
	}

	if fsc != nil {
		searchService = searchService.FetchSourceContext(fsc)
	}

	return searchService
}
