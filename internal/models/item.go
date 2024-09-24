package models

import (
	"time"
)

type Item struct {
	ID           int
	Date         time.Time
	Sum          float32
	CategoryName string
	Description  string
}
