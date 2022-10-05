package model

import (
	_ "database/sql"
	"fmt"
	"todo/db"
	"todo/dto"

	validator "github.com/go-playground/validator/v10"
)

func GetTodo() dto.Todos {
	con := db.Connect()
	sqlStatement := "SELECT * FROM todos"
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
func PostTodos(name string) error {
	con := db.Connect()
	valid := validator.New()
	data := dto.Todo{
		Name: name,
	}
	valid.Struct(data)
	sqlStatement := "INSERT into todos (name) VALUES (?)"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	result, err := statement.Exec(name)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func UpdateTodos(id int, name string) error {
	con := db.Connect()
	valid := validator.New()
	data := dto.Todo{
		Id:   id,
		Name: name,
	}
	valid.Struct(data)
	sqlStatement := "UPDATE todos set name ? where id ?"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	result, err := statement.Exec(id, name)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func DeleteTodo(id int) error {
	return nil
}
