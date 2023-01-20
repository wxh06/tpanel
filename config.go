package tpanel

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Filename string `toml:"-"`
	CaddyBin string `toml:"caddy_bin"` // Path to Caddy binary
	GOPROXY  string `toml:"goproxy,omitempty"`
}

func ReadConfig(path string) (config Config, err error) {
	_, err = toml.DecodeFile(path, &config)
	config.Filename = path
	if config.CaddyBin == "" {
		config.CaddyBin = "./caddy"
	}
	return
}

func WriteConfig(config Config) error {
	f, err := os.OpenFile(config.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	return toml.NewEncoder(f).Encode(config)
}
