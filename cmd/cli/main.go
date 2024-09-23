package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tksasha/balance/internal/server/db"
	"github.com/tksasha/balance/internal/server/env"
	"github.com/tksasha/balance/internal/server/workdir"
)

var _ = fmt.Println //nolint:forbidigo

type Item struct {
	Date        time.Time
	Sum         float32
	Category    string
	Description string
}

func (i Item) String() string {
	return fmt.Sprintf(
		"%s, %0.02f, %s, %s",
		i.Date.Format(time.DateOnly),
		i.Sum,
		i.Category,
		i.Description,
	)
}

func main() {
	workdir, err := workdir.New()
	if err != nil {
		panic(err)
	}

	db, err := db.Open(workdir, env.Get())
	if err != nil {
		panic(err)
	}

	query := `
		SELECT
			items.date,
			items.sum,
			categories.name,
			items.description
		FROM
			items
		INNER JOIN
			categories
		ON
			categories.id=items.category_id
		WHERE
			date BETWEEN "2024-09-01" AND "2024-09-30"
		AND
			items.currency=0
		ORDER BY
			items.date DESC
	`

	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	var items []Item

	for rows.Next() {
		var item Item

		if err := rows.Scan(&item.Date, &item.Sum, &item.Category, &item.Description); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Printf("%s\n", item) //nolint:forbidigo
	}
}
