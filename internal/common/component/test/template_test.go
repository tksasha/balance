package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/common/component"
	"gotest.tools/v3/assert"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TestTemplate(t *testing.T) {
	component := component.New()

	w := bytes.NewBuffer([]byte{})

	if err := component.Template(Tr(Td(Text("Name")))).Render(w); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, w.String(), "<template><tr><td>Name</td></tr></template>")
}
