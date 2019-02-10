package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "ipkiss"
	app.Usage = "mask specific columns from CSV"
	app.Version = "0.0.1"
	app.Run(os.Args)
}
