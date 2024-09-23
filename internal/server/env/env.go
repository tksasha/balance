package env

import (
	"os"
)

const (
	ENVVAR  = "BALANCE_ENV"
	DEFAULT = "development"
)

func Get() string {
	env := os.Getenv(ENVVAR)

	if env != "" {
		return env
	}

	return DEFAULT
}
