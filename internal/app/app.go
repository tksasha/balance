package app

import (
	"database/sql"
)

type App struct {
	DB *sql.DB
}

func New(db *sql.DB) *App {
	return &App{
		DB: db,
	}
}
