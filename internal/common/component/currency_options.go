package component

import (
	"maps"
	"slices"
	"strings"

	"github.com/tksasha/balance/internal/common/currency"
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) CurrencyOptions(selected string) Node {
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
