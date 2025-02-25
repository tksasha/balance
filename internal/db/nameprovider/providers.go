package nameprovider

import (
	"path/filepath"

	"github.com/tksasha/balance/internal/db/env"
	"github.com/tksasha/balance/internal/db/workdir"
)

type NameProvider struct {
	env string
}

func New() *NameProvider {
	return &NameProvider{
		env: env.Get(),
	}
}

func (p *NameProvider) Provide() string {
	return filepath.Join(workdir.New(), p.env+".sqlite3")
}

type TestNameProvider struct{}

func NewTestProvider() *TestNameProvider {
	return &TestNameProvider{}
}

func (p *TestNameProvider) Provide() string {
	return filepath.Join(workdir.New(), "test.sqlite3")
}
