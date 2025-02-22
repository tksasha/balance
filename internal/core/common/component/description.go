package component

import "regexp"

// Description parse tags from description and makes it HTML
// `[tag] description` -> `<div class="tag">tag</div> description`
func (c *Component) Description(description string) string {
	mask := regexp.MustCompile(`\[([\p{L}\s\'\-\.]+)\]`)

	return mask.ReplaceAllString(description, `<div class="tag">$1</div>`)
}
