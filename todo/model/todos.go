package model

import (
	_ "database/sql"
	"fmt"
	"net/http"
	"todo/db"
	"todo/dto"

	validator "github.com/go-playground/validator/v10"
)

func GetTodo() dto.Todos {
	con := db.Connect()
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
func PostTodos(name string) (dto.Response, error) {
	var res dto.Response
	con := db.Connect()
	valid := validator.New()
	data := dto.Todo{
		Name: name,
	}
	valid.Struct(data)
	sqlStatement := "INSERT into todos (name) VALUES (?)"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer statement.Close()
	result, err := statement.Exec(name)
	if err != nil {
		return res, err
	}
	last_id, err := result.LastInsertId()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"id_inserted": last_id,
	}
	return res, nil
}

func UpdateTodos(name string, id int) (dto.Response, error) {
	var res dto.Response
	con := db.Connect()
	valid := validator.New()
	data := dto.Todo{
		Id:   id,
		Name: name,
	}
	valid.Struct(data)
	sqlStatement := "UPDATE todos SET name = ? WHERE id = ?"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer statement.Close()
	result, err := statement.Exec(name, id)
	if err != nil {
		return res, err
	}
	row_affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"row_affected": row_affected,
	}
	fmt.Println(result)
	return res, nil
}

func DeleteTodo(id int) (dto.Response, error) {
	var res dto.Response
	con := db.Connect()
	sqlStatement := "DELETE FROM todos WHERE id = ?"
	statement, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer statement.Close()
	result, err := statement.Exec(id)
	if err != nil {
		return res, err
	}
	changed, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	row_affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"row_affected": row_affected,
	}
	fmt.Println(changed)
	return res, err
}
