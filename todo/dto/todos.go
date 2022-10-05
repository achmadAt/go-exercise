package dto

type Todos struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	IsDone bool   `json:"is_done,omitempty"`
}
