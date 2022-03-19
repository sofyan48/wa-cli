package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func (handler *CLIMapping) Logout() cli.Command {
	command := cli.Command{}
	command.Name = "logout"
	command.Usage = "logout [option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "QR path",
			Destination: &Args.QrPath,
		},
		// cli.BoolFlag{
		// 	Name:        "decryption",
		// 	Usage:       "--decryption",
		// 	Destination: &Args.Decryption,
		// },
	}
	command.Action = func(c *cli.Context) error {
		checkEnv()
		err := os.Remove(os.TempDir() + "/whatsappSession.gob")
		if err != nil {
			fmt.Println("Logout Failed")
			os.Exit(1)
		}
		fmt.Println("Logout Success")
		return nil
	}

	return command
}
