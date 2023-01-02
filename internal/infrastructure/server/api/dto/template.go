package dto

type Template struct {
	Id     int64  `json:"id,omitempty"`
	Active bool   `json:"active"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Body   string `json:"body"`
	Title  string `json:"title"`
}
