package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli"
)

func swapFiles(fName1 string, fName2 string) error {
	if fName1 == "" || fName2 == "" {
		return nil
	}
	tmpName := "/tmp/" + fName1 // create a temporary name in /tmp
	err := os.Rename(fName1, tmpName)
	err = os.Rename(fName2, fName1)
	err = os.Rename(tmpName, fName2)
	if err != nil {
		return errors.New("swap failure")
	}
	return nil
}

func main() {
	app := &cli.App{
		Name:  "swap",
		Usage: "swap [file1] [file2]",
		Action: func(c *cli.Context) error {
			if c.NArg() > 1 {
				fName1 := c.Args().Get(0)
				fName2 := c.Args().Get(1)
				res := swapFiles(fName1, fName2)
				return res
			}
			return errors.New("Usage: swap [file1] file[2]")
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
