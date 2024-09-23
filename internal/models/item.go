package models

import (
	"time"
)

type Item struct {
	ID          int
	Date        time.Time
	Sum         float32
	Category    string
	Description string
}
