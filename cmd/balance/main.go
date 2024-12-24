package main

import (
	"github.com/tksasha/balance/internal/wire"
)

func main() {
	wire.InitializeServer().Run()
}
