package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/wxh06/tpanel"
)

func main() {
	// Configuration
	var config tpanel.Config
	if cfg, err := tpanel.ReadConfig(getConfigFile()); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	} else {
		config = cfg
	}

	if err := os.Setenv("GOPROXY", config.GOPROXY); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Handle the SIGINT and SIGTERM signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		if err := tpanel.WriteConfig(config); err != nil {
			panic(err)
		}
	}()

	// Build Caddy binary
	if _, err := os.Stat(config.CaddyBin); errors.Is(err, os.ErrNotExist) {
		if err := tpanel.BuildCaddy(ctx, config.CaddyBin); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	// Start Caddy server
	if caddy, err := tpanel.StartCaddy(ctx, config.CaddyBin); err != nil {
		panic(err)
	} else {
		defer caddy.Wait()
	}
}

// Get the path to config file from command-line arguments.
func getConfigFile() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "configs/config.toml"
}
