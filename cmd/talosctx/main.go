package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/lederniermetre/talosctx/internal/cmdutil"
	"github.com/lederniermetre/talosctx/internal/talosconfig"
	"gitlab.com/greyxor/slogor"
	"gopkg.in/yaml.v3"
)

func main() {
	var getContext bool
	var debug bool
	flag.BoolVar(&getContext, "get-context", false, "Get only current context")
	flag.BoolVar(&debug, "debug", false, "enable log level debug")
	flag.Parse()

	cmdutil.SetLogger(debug)

	yamlData, _, err := talosconfig.GetTalosConfigContent()
	if err != nil {
		slog.Error("Error when get content", slogor.Err(err))
		os.Exit(1)
	}

	if getContext {
		slog.Debug("Retrive only current context")
		var config talosconfig.Talosconfig
		err = yaml.Unmarshal(yamlData, &config)
		if err != nil {
			slog.Error("Error when get content", slogor.Err(err))
			os.Exit(1)
		}

		fmt.Printf("%+v", config.Context)
	} else {
		fmt.Printf("%s\n", yamlData)
	}
}
