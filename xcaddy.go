package tpanel

import (
	"context"

	"github.com/caddyserver/xcaddy"
)

func Build(ctx context.Context, outputFile string) error {
	builder := xcaddy.Builder{}
	return builder.Build(ctx, outputFile)
}
