package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	UserID          int
	API             string
	CssVersion      string
}

//go:embed fe_templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {

	td.API = app.config.api

	// if app.Session.Exists(r.Context(), "userID") {
	// 	// app.infoLog.Println("session userID exist")
	// 	td.IsAuthenticated = 1
	// 	td.UserID = app.Session.GetInt(r.Context(), "userID")
	// 	// app.infoLog.Println(td.UserID)
	// } else {
	// 	// app.infoLog.Println("session userID not exist")
	// 	td.IsAuthenticated = 0
	// 	td.UserID = 0
	// 	// app.infoLog.Println(td.UserID)

	// }

	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	templateToRender := fmt.Sprintf("fe_templates/%s.page.gohtml", page)

	_, templateInMap := app.templateCache[templateToRender]

	if templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("fe_templates/%s.partial.gohtml", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).ParseFS(templateFS, "fe_templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).ParseFS(templateFS, "fe_templates/base.layout.gohtml", templateToRender)

	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
