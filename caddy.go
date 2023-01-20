package tpanel

import (
	"context"
	"os/exec"

	"github.com/caddyserver/xcaddy"
)

func BuildCaddy(ctx context.Context, outputFile string) error {
	builder := xcaddy.Builder{}
	return builder.Build(ctx, outputFile)
}

func StartCaddy(ctx context.Context, caddyBin string) (*exec.Cmd, error) {
	caddy := exec.CommandContext(ctx, caddyBin, "run")
	return caddy, caddy.Start()
}
