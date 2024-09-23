package workdir

import (
	"os"
)

const NAME = ".balance"

func New() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	workDir := userHomeDir + string(os.PathSeparator) + NAME

	if err := os.MkdirAll(workDir, 0o750); err != nil { //nolint:mnd
		return "", err
	}

	return workDir, nil
}
