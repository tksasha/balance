package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/item/component"
	"gotest.tools/v3/assert"
)

func TestDescription(t *testing.T) {
	testmap := map[string]string{
		"":                             ``,
		"some description":             `some description`,
		"[superstore] purchases":       `<div class="tag">superstore</div> purchases`,
		"[First Tag] some description": `<div class="tag">First Tag</div> some description`,
		"[імперія м'яса] ковбаски":       `<div class="tag">імперія м'яса</div> ковбаски`,
		"[а-банк] відсотки за депозитом": `<div class="tag">а-банк</div> відсотки за депозитом`,
		"[flowers.ua] доставка квітів":   `<div class="tag">flowers.ua</div> доставка квітів`,
	}

	testmap["[First Tag] [Second Tag] some description"] = `` +
		`<div class="tag">First Tag</div> <div class="tag">Second Tag</div> some description`

	component := component.New()

	for description, expected := range testmap {
		w := bytes.NewBuffer([]byte{})

		if err := component.Description(description).Render(w); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, w.String(), expected)
	}
}
