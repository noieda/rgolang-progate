package main

import (
	"net/http"
	"rgolang-progate/internal/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) ViewTodolist(w http.ResponseWriter, r *http.Request) {

	// app.infoLog.Printf("disini")

	TodoList, err := app.DB.ViewTodolist()
	if err != nil {
		app.badRequest(w, r, err)
		app.errorLog.Println(err)
	}

	// app.infoLog.Println(TodoList)

	app.writeJSON(w, http.StatusOK, TodoList)
}

func (app *application) ViewTodo(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	todoID, _ := strconv.Atoi(id)

	if todoID == 0 {
		return
	}

	todo, err := app.DB.ViewTodo(todoID)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, todo)
}

func (app *application) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	todoID, _ := strconv.Atoi(id)

	app.infoLog.Println("update to do => ", id)

	var todo models.Todolist

	err := app.readJSON(w, r, &todo)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	if todoID > 0 {

		app.infoLog.Println("update user dgn id => ", id)

		err = app.DB.UpdateTodo(todo)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

	} else {

		app.infoLog.Println("create user baru")

		err = app.DB.CreateTodo(todo)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	app.writeJSON(w, http.StatusOK, resp)
}

func (app *application) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	todoID, _ := strconv.Atoi(id)

	app.infoLog.Println("delete to do => ", id)
	err := app.DB.DeleteTodo(todoID)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	app.writeJSON(w, http.StatusOK, resp)
}
