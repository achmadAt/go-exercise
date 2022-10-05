package entities

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	Users = map[int]*User{}
	Seq   = 1
)
