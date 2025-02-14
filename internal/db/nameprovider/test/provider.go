package nameprovider

import (
	"fmt"
	"os"

	"github.com/tksasha/balance/internal/workdir"
)

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Provide() string {
	return fmt.Sprintf(
		"%s%s%s.sqlite3?parseTime=true",
		workdir.New(),
		string(os.PathSeparator),
		"test",
	)
}
