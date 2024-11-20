package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/bdronneau/talosctx/internal/cmdutil"
	"github.com/bdronneau/talosctx/internal/printer"
	"github.com/bdronneau/talosctx/internal/talosconfig"
	"gopkg.in/yaml.v3"
)

func main() {
	var getContext bool
	var debug bool
	flag.BoolVar(&getContext, "get-context", false, "Get only current context")
	flag.BoolVar(&debug, "debug", false, "enable log level debug")
	flag.Parse()

	cmdutil.SetLogger(debug)

	yamlData, path, err := talosconfig.GetTalosConfigContent()
	if err != nil {
		slog.Error("Error when get content", slog.Any("error", err))
		os.Exit(1)
	}

	var config talosconfig.Talosconfig
	err = yaml.Unmarshal(yamlData, &config)

	if getContext {
		slog.Debug("Retrive only current context")
		if err != nil {
			slog.Error("Error when get content", slog.Any("error", err))
			os.Exit(1)
		}

		fmt.Printf("%s", config.Context)
	} else {
		context, err := printer.SelectContext(config.Contexts)
		if err != nil {
			slog.Error("Can not select context", slog.Any("error", err))
			os.Exit(1)
		}

		config.Context = context

		data, err := yaml.Marshal(&config)
		if err != nil {
			slog.Error("Can not marshal yamlData", slog.Any("error", err))
			os.Exit(1)
		}

		err = talosconfig.WritetalosConfig(path, data)
		if err != nil {
			slog.Error("Can not write", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
