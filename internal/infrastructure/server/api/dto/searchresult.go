package dto

type SearchResult struct {
	Items []Template `json:"items"`
	Pages int        `json:"pages"`
	Total int        `json:"total"`
}
