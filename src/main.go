package main

import (
	"os"

	"github.com/orn/wa-cli/src/cmd"
)

func main() {
	app := cmd.AppCommands()
	app.Run(os.Args)
}
