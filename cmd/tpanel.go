package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

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

	caddy := exec.Command(config.CaddyBin, "run")

	// Handle the SIGINT and SIGTERM signal
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		if err := tpanel.WriteConfig(config); err != nil {
			panic(err)
		}
		if err := caddy.Process.Signal(syscall.SIGTERM); err != nil {
			panic(err)
		}
	}()

	// Run Caddy server
	if err := caddy.Run(); err != nil {
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
