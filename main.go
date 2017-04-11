package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "got"
	app.Usage = "Tool for repository support on import paths."
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello world!")
		return nil
	}
	app.Run(os.Args)
}
