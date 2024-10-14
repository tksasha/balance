package main

import (
	"github.com/tksasha/balance/internal/server"
)

func main() {
	server.
		Initialize().
		Run()
}
