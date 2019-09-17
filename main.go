package main


import (
	"fmt"
	"log"
	"main/cli_output"
	"main/images"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "DcontainerTkt"
	app.Version = "0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Dong Vo Thanh",
			Email: "vothanhdong18@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "images",
			Usage: "list all images",
			Action: func(c *cli.Context) error {
				images, _ := images.ListContainerImages()
				cli_output.PrintListImage(images)
				return nil
			},
		},
		{
			Name:  "image",
			Usage: "action with images",
			Subcommands: []cli.Command{
				{
					Name:  "ls",
					Usage: "list all images",
					Action: func(c *cli.Context) error {
						images, _ := images.ListContainerImages()
						cli_output.PrintListImage(images)
						return nil
					},
				},
				{
					Name:  "rm",
					Usage: "remove image",
					Action: func(c *cli.Context) error {
						err := images.RemoveContainerImages(c.Args().First())
						if err != nil {
							fmt.Println(err)
							return nil

						}

						fmt.Println("Image was removed: " + c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name:  "pull",
			Usage: "pull image",
			Action: func(c *cli.Context) error {
				images.PullContainerImage(strings.Split(c.Args().First(), ":")[0], strings.Split(c.Args().First(), ":")[1])
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


