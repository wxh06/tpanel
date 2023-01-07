package main

import (
	"context"
	"errors"
	"os"

	"github.com/wxh06/tpanel"
)

func main() {
	config := tpanel.ReadConfig(getConfigFile())

	if err := os.Setenv("GOPROXY", config.GOPROXY); err != nil {
		panic(err)
	}

	if _, err := os.Stat(config.CaddyBin); errors.Is(err, os.ErrNotExist) {
		if err := tpanel.Build(context.TODO(), config.CaddyBin); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	if err := tpanel.WriteConfig(config); err != nil {
		panic(err)
	}
}

// Get the path to config file from command-line arguments.
func getConfigFile() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "config.toml"
}
