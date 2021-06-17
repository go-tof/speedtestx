package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-tof/speedtestx/cmd/speedtestx/speedtest"
	"github.com/urfave/cli/v2"
)

// Version set at compile-time.
const Version = "1.0.0"

func main() {
	// Cli application instance.
	app := &cli.App{
		Name:      "speedtestx",
		Usage:     "based speedtest check network bandwidth cli application",
		Copyright: "Copyright (c) " + strconv.Itoa(time.Now().Year()) + " helloshaohua, contact email address wu.shaohua@foxmail.com",
		Version:   Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "all",
				Usage: "All check network bandwidth result data.",
			},
			&cli.StringFlag{
				Name:  "upload",
				Usage: "All check network upload bandwidth result data.",
			},
			&cli.StringFlag{
				Name:  "download",
				Usage: "All check network download bandwidth result data.",
			},
		},
		Action: speedtest.ActionFuncForSpeedtest,
	}

	// Run cli application.
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
