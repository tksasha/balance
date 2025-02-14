package nameprovider

import (
	"fmt"
	"os"

	"github.com/tksasha/balance/internal/db/env"
	"github.com/tksasha/balance/internal/db/workdir"
)

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Provide() string {
	return fmt.Sprintf(
		"%s%s%s.sqlite3",
		workdir.New(),
		string(os.PathSeparator),
		env.Get(),
	)
}
