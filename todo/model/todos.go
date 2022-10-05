package model

import (
	_ "database/sql"
	"fmt"
	"todo/config"
	"todo/dto"
)

func GetTodo() dto.Todos {
	con := config.Connect()
	sqlStatement := "SELECT id, name FROM todos"
	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer rows.Close()
	result := dto.Todos{}
	for rows.Next() {
		todo := dto.Todo{}
		err := rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		result.Todos = append(result.Todos, todo)
	}
	return result
}
