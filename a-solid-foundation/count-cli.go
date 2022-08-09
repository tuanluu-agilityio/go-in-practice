package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down."
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count Up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Value: 10,
					Usage: "Value to count up to",
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start < 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 1; i < start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "downn",
			ShortName: "d",
			Usage:     "Count Down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Value: 10,
					Usage: "Start counting down from",
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("Start cannot be negative.")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
