package main

import (
	"context"

	"github.com/fremantle-industries/workshop/cmd"
)

func main() {
	cmd.AddCommands()
	cmd.Execute(context.Background())
}
