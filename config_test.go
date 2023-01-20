package tpanel

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	config, err := ReadConfig(f.Name())
	if err != nil {
		t.Error(err)
	}
	if config.Filename != f.Name() {
		t.Error("Incorrect file path in configuration")
	}

	err = WriteConfig(config)
	if err != nil {
		t.Error(err)
	}
}
