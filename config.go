package tpanel

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Filename string `toml:"-"`
	CaddyBin string `toml:"caddy_bin"`
	GOPROXY  string `toml:"goproxy,omitempty"`
}

func ReadConfig(path string) (config Config) {
	config.Filename = path
	_, err := toml.DecodeFile(path, &config)
	if config.CaddyBin == "" {
		config.CaddyBin = "./caddy"
	}
	if errors.Is(err, os.ErrNotExist) {
		if err := WriteConfig(config); err != nil {
			panic(err)
		}
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
