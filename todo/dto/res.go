package dto

type Response struct {
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
