package workdir

import (
	"os"
)

const NAME = ".balance"

func New() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	workdir := userHomeDir + string(os.PathSeparator) + NAME

	if err := os.MkdirAll(workdir, 0o750); err != nil { //nolint:mnd
		panic(err)
	}

	return workdir
}
