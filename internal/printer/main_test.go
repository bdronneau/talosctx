package printer

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/bdronneau/talosctx/internal/talosconfig"
	"github.com/stretchr/testify/assert"
)

func TestSelectContext(t *testing.T) {
	go func() {
		err := keyboard.SimulateKeyPress(keys.Down)
		assert.NoError(t, err)
		err = keyboard.SimulateKeyPress(keys.Enter)
		assert.NoError(t, err)
	}()

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

	context, err := SelectContext(data)
	assert.NoError(t, err)
	assert.Equal(t, "2", context)
}
