package models

import "example/db"

func Show(id string) (todo Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer connection.Close()

	err = connection.QueryRow(`SELECT * FROM todos WHERE ID = $1`, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return todo, err
}
