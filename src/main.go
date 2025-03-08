package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the current version",
	}

	app := &cli.App{
		Name:     "hermes",
		Usage:    "A lightweight, high-performance message queue system built in Go.",
		Version:  "v1.0.0",
		Commands: commands(),
		Flags:    flags(),
		Action:   action,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func commands() []*cli.Command {
	return []*cli.Command{
		{
			// TODO: is this really needed? Since the server will run in a container...
			Name:    "logs",
			Aliases: []string{"l"},
			Usage:   "print past logs",
			Action: func(cCtx *cli.Context) error {
				fmt.Println("boom! here are the logs!")
				// Call the main function that will show past logs stored inside the container
				return nil
			},
		},
		{
			Name:    "subscribe",
			Aliases: []string{"s"},
			Usage:   "runs the subscriber server while subscribing to a topic",
			Action: func(ctx *cli.Context) error {
				// Ask the user for the topic, use a sub command or a sub flag to get the topic
				fmt.Println("subscribing to a topic here")
				return nil
			},
		},
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "publish your message for a topic",
			Action: func(ctx *cli.Context) error {
				// Ask the user for the topic and the message to publish
				fmt.Println("publishing a new message to a specific topic")
				return nil
			},
		},
	}
}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "run while printing all the logs",
			Aliases: []string{"l"},
		},
	}
}

func action(cCtx *cli.Context) error {
	fmt.Println("Hello")
	// Call the main server operation, should start the broker here
	return nil
}
