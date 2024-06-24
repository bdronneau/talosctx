package printer

import (
	"github.com/bdronneau/talosctx/internal/talosconfig"
	"github.com/pterm/pterm"
)

func SelectContext(data map[string]talosconfig.Taloscontext) (string, error) {
	var entries []string

	for entry := range data {
		entries = append(entries, entry)
	}

	selectedOption, err := pterm.DefaultInteractiveSelect.WithOptions(entries).Show()
	if err != nil {
		return "", err
	}

	pterm.Info.Printfln("Selected context: %s", pterm.Green(selectedOption))

	return selectedOption, nil
}
