package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"rgolang-progate/internal/driver"
	"rgolang-progate/internal/models"
	"time"
)

const version = "1.0.0"
const cssVersion = "1"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	secretkey string
	frontend  string
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	DB            models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Println(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "trevor:secret@tcp(localhost:3306)/todolist?parseTime=true&tls=false", "DSN")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "url to frontend")
	flag.StringVar(&cfg.api, "api", "http://localhost:4002", "URL to api")

	infoLog := log.New(os.Stdout, "INFO frontend\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR frontend\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	tc := make(map[string]*template.Template)
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		DB:            models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
