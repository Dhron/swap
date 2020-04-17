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
				file1 := c.Args().Get(0)
				file2 := c.Args().Get(1)
				tmpName := "/tmp/" + file1 // create a temporary name
				os.Rename(file1, tmpName)
				os.Rename(file2, file1)
				os.Rename(tmpName, file2)
				fmt.Println("swapped " + file1 + " with " + file2)
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
