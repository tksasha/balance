package providers

import (
	"fmt"
	"os"

	"github.com/tksasha/balance/internal/workdir"
)

type DBNameProvider struct{}

func NewDBNameProvider() *DBNameProvider {
	return &DBNameProvider{}
}

func (p *DBNameProvider) Provide() string {
	return fmt.Sprintf(
		"%s%s%s.sqlite3",
		workdir.New(),
		string(os.PathSeparator),
		"test",
	)
}
