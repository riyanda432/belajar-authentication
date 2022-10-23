package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	command_unit_test "github.com/riyanda432/belajar-authentication/src/interface/cli/command/unit_test"
	cli "github.com/urfave/cli/v2"
)

var uTExclude = command_unit_test.UTExclude

func main() {

	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name:  "unit-test-validation-with-excluded-folder",
			Usage: "run unit test with ignored package",
			Action: func(c *cli.Context) error {
				return uTExclude()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
