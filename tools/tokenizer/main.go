package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/takuoki/gsheets"
	"github.com/urfave/cli"
)

const version = "1.0.0"

func main() {

	app := cli.NewApp()
	app.Name = "googleAPITokenizer"
	app.Version = version
	app.Usage = "This tool generates a token for Google Sheets API."
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "gen",
			Usage: "Generate token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "credentials, c",
					Value: "credentials.json",
					Usage: "credentials file name",
				},
				cli.BoolFlag{
					Name:  "writable, w",
					Usage: "whether to create a writable token",
				},
			},
			Action: func(c *cli.Context) error {
				if _, err := os.Stat("token.json"); err == nil {
					return errors.New("'token.json' file already exists. if you want to regenerate, remove it and try again")
				}
				if c.String("credentials") == "" {
					return errors.New("credentials file name is empty")
				}
				if c.Bool("writable") {
					if _, err := gsheets.NewForCLI(context.Background(), c.String("credentials"), gsheets.ClientWritable()); err != nil {
						return err
					}
				} else {
					if _, err := gsheets.NewForCLI(context.Background(), c.String("credentials")); err != nil {
						return err
					}
				}
				fmt.Println("complete")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		return
	}
}
