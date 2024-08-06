package printer

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/bdronneau/talosctx/internal/talosconfig"
	"github.com/stretchr/testify/assert"
)

func TestSelectContext(t *testing.T) {
	data := map[string]talosconfig.Taloscontext{
		"1": {
			Ca:  "dog",
			Crt: "cat",
			Key: "human",
		},
		"2": {
			Ca:  "cow",
			Crt: "fish",
			Key: "human",
		},
		"3": {
			Ca:  "cow",
			Crt: "fish",
			Key: "human",
		},
	}

	go func() {
		_ = keyboard.SimulateKeyPress(keys.Down)
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	context, err := SelectContext(data)
	assert.NoError(t, err)
	assert.Equal(t, "2", context)
}
