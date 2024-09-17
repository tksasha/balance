package main

import (
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/server"
)

func main() {
	server.Run(
		config.New(),
	)
}
