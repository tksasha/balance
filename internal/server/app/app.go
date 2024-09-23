package app

import (
	"database/sql"
	"html/template"
)

type App struct {
	T  *template.Template
	DB *sql.DB
}
