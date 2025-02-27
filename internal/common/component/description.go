package component

import (
	"regexp"

	. "maragu.dev/gomponents" //nolint:stylecheck
)

// Description parse tags from description and makes it HTML
// `[tag] description` -> `<div class="tag">tag</div> description`
func (c *Component) Description(description string) Node {
	mask := regexp.MustCompile(`\[([\p{L}\s\'\-\.]+)\]`)

	description = mask.ReplaceAllString(description, `<div class="tag">$1</div>`)

	return Raw(description)
}
