package app

import (
	"database/sql"
)

type App struct {
	DB *sql.DB
}
