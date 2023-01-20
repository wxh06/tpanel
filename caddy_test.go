package tpanel

import (
	"context"
	"os"
	"testing"
)

func TestCaddy(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	ctx, cancel := context.WithCancel(context.Background())

	if err := BuildCaddy(ctx, f.Name()); err != nil {
		t.Error(err)
	}

	if _, err := StartCaddy(ctx, f.Name()); err != nil {
		t.Error(err)
	}
	cancel()
}
