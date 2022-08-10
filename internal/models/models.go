package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// DBModel is type for database connection values
type DBModel struct {
	DB *sql.DB
}

// Models is wrapper for all models
type Models struct {
	DB DBModel
}

type Todolist struct {
	ID         int       `json:"id"`
	Todo       string    `json:"todo"`
	PoC        string    `json:"poc"`
	Deadline   time.Time `json:"deadline"`
	TodoStatus string    `json:"status"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

func (m *DBModel) ViewTodolist() ([]*Todolist, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var todos []*Todolist

	query := `
		select
			id, todo, poc, deadline, todo_status
		from
			todolist
		order by
			todo
	`

	// fmt.Printf("111")

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// fmt.Printf("222")

	for rows.Next() {

		var todo Todolist
		err = rows.Scan(
			&todo.ID,
			&todo.Todo,
			&todo.PoC,
			&todo.Deadline,
			&todo.TodoStatus,
			// &todo.CreatedAt,
			// &todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}

func (m *DBModel) ViewTodo(id int) (Todolist, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var todo Todolist

	row := m.DB.QueryRowContext(ctx, `
		select
			id, todo, poc, deadline, todo_status
		from
			todolist
		where id = ?
	`, id)

	err := row.Scan(
		&todo.ID,
		&todo.Todo,
		&todo.PoC,
		&todo.Deadline,
		&todo.TodoStatus,
	)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (m *DBModel) CreateTodo(todo Todolist) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// var td Todolist

	stmt := `
		insert into
			todolist
				(
					todo, poc, deadline, todo_status, created_at, updated_at
				)
			values
				(
					?, ?, ? ,?, ?, ?
				)
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		todo.Todo,
		todo.PoC,
		time.Now(),
		todo.TodoStatus,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateTodo(todo Todolist) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		update
			todolist
		set
			todo = ?,
			poc = ?,
			deadline = ?,
			todo_status = ?,
			updated_at = ?
		where
			id = ?

	`

	_, err := m.DB.ExecContext(ctx, stmt,
		todo.Todo,
		todo.PoC,
		todo.Deadline,
		todo.TodoStatus,
		time.Now(),
		todo.ID,
	)
	if err != nil {
		return err
	}

	fmt.Println("updateToDo executed")

	return nil
}

func (m *DBModel) DeleteTodo(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		delete from
			todolist
		where
			id = ?
	`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
