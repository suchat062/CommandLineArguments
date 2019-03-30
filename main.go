package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "start",
			Usage: "Start Date",
		},
		cli.StringFlag{
			Name: "end",
			Usage: "End Date",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Format Parameter (YYYY-MM-DD) => -start=0000-00-00 -end=0000-00-00")
		fmt.Println()
		Start := c.String("start")
		End := c.String("end")
		tStart, err := time.Parse("2006-01-02", Start)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		tEnd, err := time.Parse("2006-01-02", End)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if tStart.Unix() <= tEnd.Unix() {
			fmt.Println(Start)
			fmt.Println(tStart.Unix())

			fmt.Println(End)
			fmt.Println(tEnd.Unix())
		}else {
			fmt.Println("Start Date And End Date Error")
			os.Exit(1)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}