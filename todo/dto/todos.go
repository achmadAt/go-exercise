package dto

type Todo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Todos struct {
	Todos []Todo `json:"todos,omitempty"`
}
