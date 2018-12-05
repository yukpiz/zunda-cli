package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zunda"
	app.Usage = "Zundamaru CLI tool."
	app.Version = "0.0.1"

	app.Action = Handler

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "ariaki, a",
			Usage: "Spawn ariaki",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(context *cli.Context) error {
	if context.Bool("ariaki") {
		return ariaki()
	}
	return zunda()
}

func zunda() error {
	data, err := Asset("zunda.txt")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(data))
	return nil
}

func ariaki() error {
	data, err := Asset("ariaki.txt")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(data))
	return nil
}
