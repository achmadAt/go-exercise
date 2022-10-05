package model

import (
	_ "database/sql"
	"fmt"
	"todo/config"
	//"todo/dto"
)

type Todo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Todos struct {
	Todos []Todo `json:"todos,omitempty"`
}

func GetTodo() Todos {
	con := config.Connect()
	sqlStatement := "SELECT id, name FROM todos"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer rows.Close()
	result := Todos{}
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		result.Todos = append(result.Todos, todo)
	}
	return result
}
