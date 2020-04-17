// main.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "swap",
		Usage: "swap [file1] [file2]",
		Action: func(c *cli.Context) error {
			if c.NArg() > 1 {
				fName1 := c.Args().Get(0)
				fName2 := c.Args().Get(1)
				tmpName := "/tmp/" + fName1 // create a temporary name
				os.Rename(fName1, tmpName)
				os.Rename(fName2, fName1)
				os.Rename(tmpName, fName2)
				fmt.Println("swapped " + fName1 + " with " + fName2)
				return nil
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
