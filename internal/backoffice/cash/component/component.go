package component

import (
	"maps"
	"slices"
	"strings"

	"github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/common/currency"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type Component struct {
	*component.Component
}

func New() *Component {
	return &Component{
		Component: component.New(),
	}
}

func (c *Component) currencyOptions(selected string) Node {
	options := []Node{}

	currencies := currency.All()

	keys := slices.Collect(maps.Keys(currencies))

	slices.Sort(keys)

	for _, currency := range keys {
		code := currencies[currency]

		codeToLower := strings.ToLower(code)

		option := Option(
			Value(codeToLower),
			Label(Text(code)),
			If(codeToLower == selected, Selected()),
		)

		options = append(options, option)
	}

	return Group(options)
}
