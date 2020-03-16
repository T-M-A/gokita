package main

import (
	"log"
	"os"

	"github.com/SeamPay/gokita/cmd/service"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "gokita",
		Description: "SeamPay API microservice boilerplate",
		Usage:       "SeamPay API microservice boilerplate",
		UsageText:   "gokita [global options] command [command options] [arguments...]",
		Authors: []*cli.Author{
			{
				Name:  "SeamPay",
				Email: "hello@seampay.com",
			},
		},
		Copyright: "(c) 2020 SeamPay Limited",
		Commands:  service.Commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
