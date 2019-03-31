package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	var start, end string
	start = "Command Line Interface"
	end = "Command Line Interface"

	app := cli.NewApp()
	app.Name = "Command Line Claim"
	app.Usage = "By Allianz.com (NENG)"
	app.UsageText = "Date Format Parameter (YYYY-MM-DD) => -start=0000-00-00 -end=0000-00-00"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "suchat rungchaeng",
			Email: "suchat.ru@allianz.com",
		},
	}
	app.Copyright = "(c) 2019 allianz.com"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "start, s",
			Usage: "Start Date Format Parameter (YYYY-MM-DD) => -start=0000-00-00",
			Destination: &start,
		},
		cli.StringFlag{
			Name: "end, e",
			Usage: "End Date Format Parameter (YYYY-MM-DD) => -end=0000-00-00",
			Destination: &end,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Format Parameter (YYYY-MM-DD) => -start=0000-00-00 -end=0000-00-00")
		fmt.Println()
		Start := c.String("start")
		End := c.String("end")
		tStart, err := time.Parse("2006-01-02", Start)
		if err != nil {
			//fmt.Println(err.Error())
			return cli.NewExitError("Error => Start Date not format, -start=0000-00-00", 86)
			//os.Exit(1)
		}

		tEnd, err := time.Parse("2006-01-02", End)
		if err != nil {
			//fmt.Println(err.Error())
			return cli.NewExitError("Error => End Date not format, -end=0000-00-00", 86)
			//os.Exit(1)
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