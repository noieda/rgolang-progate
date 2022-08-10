package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) ViewTodo(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "todo", &templateData{
		// Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
