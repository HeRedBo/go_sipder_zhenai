package model

type SearchResult struct {
	Hits int64
	Start int
	Query string
	TotalPage int64
	CurrentPage int
	Items []interface{}
}
