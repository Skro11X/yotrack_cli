package main

import (
	"context"
	"os"

	"try_parse_youtrack/internal/cli"
)

func main() {
	os.Exit(cli.Run(context.Background(), os.Args[1:], os.Stdout, os.Stderr))
}
